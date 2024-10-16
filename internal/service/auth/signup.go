package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/auth"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type SignUpParams struct {
	Email        string
	Phone        string
	Password     string
	FirstName    string
	LastName     string
	MiddleName   string
	AvatarURL    string
	INN          string
	IsContractor bool
	IPAddress    string
}

type SignUpResult struct {
	User         models.User
	Organization models.Organization
	Session      models.Session
	AccessToken  string
}

func (s *Service) SignUp(ctx context.Context, params SignUpParams) (SignUpResult, error) {
	var resp SignUpResult

	organization, err := s.dadataGateway.GetOrganization(ctx, params.INN)
	if err != nil {
		return SignUpResult{}, fmt.Errorf("get organization: %w", err)
	}

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		organization, err = s.organizationStore.Create(ctx, qe, organization)
		if err != nil {
			return fmt.Errorf("create organization: %w", err)
		}

		resp.Organization = organization

		hashedPassword, err := crypto.Password(params.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		user := models.User{
			OrganizationID: organization.ID,
			Email:          params.Email,
			Phone:          params.Phone,
			PasswordHash:   hashedPassword,
			TOTPSalt:       uuid.New().String(),
			FirstName:      params.FirstName,
			LastName:       params.LastName,
			MiddleName:     params.MiddleName,
			AvatarURL:      params.AvatarURL,
			IsContractor:   params.IsContractor,
		}

		user, err = s.userStore.Create(ctx, qe, user)
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		resp.User = user

		accessToken, err := s.tokenAuthorizer.GenerateToken(auth.Payload{
			ID: user.ID,
		})
		if err != nil {
			return fmt.Errorf("generate access token: %w", err)
		}

		resp.AccessToken = accessToken

		session := models.Session{
			RefreshToken: uuid.New(),
			UserID:       user.ID,
			IPAddress:    params.IPAddress,
			ExpiresAt:    time.Now().Add(s.tokenAuthorizer.GetRefreshTokenDurationLifetime()),
		}

		session, err = s.sessionStore.Create(ctx, qe, session)
		if err != nil {
			return fmt.Errorf("create session: %w", err)
		}

		resp.Session = session

		return nil
	})
	if err != nil {
		return SignUpResult{}, fmt.Errorf("run transaction: %w", err)
	}

	return resp, nil
}
