package admin

import (
	"context"
	"fmt"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type SignInParams struct {
	Email     string
	Password  string
	IPAddress string
}

type SignInResult struct {
	User        models.AdminUser
	Session     models.Session
	AccessToken string
}

func (s *Service) SignIn(ctx context.Context, params SignInParams) (SignInResult, error) {
	user, err := s.adminStore.GetUser(ctx, s.psql.DB(), store.AdminGetUserParams{Email: params.Email})
	if err != nil {
		return SignInResult{}, fmt.Errorf("get user: %w", err)
	}

	err = crypto.CheckPassword(params.Password, user.PasswordHash)
	if err != nil {
		return SignInResult{}, cerr.Wrap(err, cerr.CodeInvalidCredentials, "invalid email or password", nil)
	}

	session, err := s.adminStore.CreateSession(ctx, s.psql.DB(), store.SessionCreateParams{
		ID:        randSessionID(sessionLength),
		UserID:    user.ID,
		IPAddress: params.IPAddress,
		ExpiresAt: time.Now().Add(s.tokenAuthorizer.GetRefreshTokenDurationLifetime()),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("create session: %w", err)
	}

	rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
		UserID: user.ID,
		Role:   int(user.Role),
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
