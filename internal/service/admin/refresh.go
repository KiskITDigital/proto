package admin

import (
	"context"
	"fmt"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Refresh(ctx context.Context, sessionToken string) (SignInResult, error) {
	session, err := s.adminStore.GetSession(ctx, s.psql.DB(), sessionToken)
	if err != nil {
		return SignInResult{}, fmt.Errorf("get session: %w", err)
	}

	user, err := s.adminStore.GetUser(ctx, s.psql.DB(), store.AdminGetUserParams{
		ID: session.UserID,
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("get user: %w", err)
	}

	rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
		UserID: user.ID,
		Role:   int(user.Role),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("generate access token: %w", err)
	}

	session, err = s.adminStore.UpdateSession(ctx, s.psql.DB(), store.SessionUpdateParams{
		ID:        session.ID,
		ExpiresAt: time.Now().Add(s.tokenAuthorizer.GetRefreshTokenDurationLifetime()),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("update session: %w", err)
	}

	return SignInResult{
		User:        user,
		Session:     session,
		AccessToken: rawToken,
	}, nil
}
