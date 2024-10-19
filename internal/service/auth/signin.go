package auth

import (
	"context"
	"fmt"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"golang.org/x/crypto/bcrypt"
)

type SignInParams struct {
	Email     string
	Password  string
	IPAddress string
}

type SignInResult struct {
	User        models.User
	Session     models.Session
	AccessToken string
}

func (s *Service) SignIn(ctx context.Context, params SignInParams) (SignInResult, error) {
	user, err := s.userStore.GetWithOrganiztion(ctx, s.psql.DB(), store.UserGetParams{Email: params.Email})
	if err != nil {
		return SignInResult{}, fmt.Errorf("get user with organization: %w", err)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password)); err != nil {
		cerr.Wrap(err, cerr.CodeInvalidCredentials, "invalid email or password", nil)
	}

	session, err := s.sessionStore.Create(ctx, s.psql.DB(), store.SessionCreateParams{
		ID:        randSessionID(sessionLength),
		UserID:    user.ID,
		IPAddress: params.IPAddress,
		ExpiresAt: time.Now().Add(RefreshTokenLifetime),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("create session: %w", err)
	}

	rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
		UserID:         user.ID,
		OrganizationID: user.Organization.ID,
		Role:           int(user.Role),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("generate access token: %w", err)
	}

	return SignInResult{
		User:        user,
		Session:     session,
		AccessToken: rawToken,
	}, nil
}
