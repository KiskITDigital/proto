package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
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
	User        models.User
	Session     models.Session
	AccessToken string
}

func (s *Service) SignUp(ctx context.Context, params SignUpParams) (SignUpResult, error) {
	var result SignUpResult

	resp, err := s.dadataGateway.FindByINN(ctx, params.INN)
	if err != nil {
		return SignUpResult{}, fmt.Errorf("get organization: %w", err)
	}

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		// FIXME: panic
		organization, err := s.organizationStore.Create(ctx, qe, store.OrganizationCreateParams{
			BrandName: resp.Suggestions[0].Data.Name.Short,
			FullName:  resp.Suggestions[0].Data.Name.FullWithOpf,
			ShortName: resp.Suggestions[0].Data.Name.ShortWithOpf,
			INN:       resp.Suggestions[0].Data.INN,
			OKPO:      resp.Suggestions[0].Data.OKPO,
			OGRN:      resp.Suggestions[0].Data.OGRN,
			KPP:       resp.Suggestions[0].Data.KPP,
			TaxCode:   resp.Suggestions[0].Data.Address.Data.TaxOffice,
			Address:   resp.Suggestions[0].Data.Address.UnrestrictedValue,
		})
		if err != nil {
			return fmt.Errorf("create organization: %w", err)
		}

		hashedPassword, err := crypto.Password(params.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		user, err := s.userStore.Create(ctx, qe, store.UserCreateParams{
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
		})
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		user.Organization = organization
		result.User = user

		session, err := s.sessionStore.Create(ctx, qe, store.SessionCreateParams{
			ID:        randSessionID(sessionLength),
			UserID:    user.ID,
			IPAddress: params.IPAddress,
			ExpiresAt: time.Now().Add(RefreshTokenLifetime),
		})
		if err != nil {
			return fmt.Errorf("create session: %w", err)
		}

		result.Session = session

		rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
			UserID:         user.ID,
			OrganizationID: user.Organization.ID,
			Role:           int(user.Role),
		})
		if err != nil {
			return fmt.Errorf("generate access token: %w", err)
		}

		result.AccessToken = rawToken

		return nil
	})
	if err != nil {
		return SignUpResult{}, fmt.Errorf("run transaction: %w", err)
	}

	return result, nil
}
