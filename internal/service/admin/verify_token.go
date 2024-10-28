package admin

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
)

func (s *Service) ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error) {
	claims, err := s.tokenAuthorizer.ValidateToken(accessToken)
	if err != nil {
		return token.Claims{}, fmt.Errorf("validate access token: %w", err)
	}

	return claims, nil
}
