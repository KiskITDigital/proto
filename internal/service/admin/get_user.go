package admin

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type GetUserParams struct {
	ID int
}

type GetUserResult struct {
	User models.AdminUser
}

func (s *Service) GetUser(ctx context.Context, params GetUserParams) (GetUserResult, error) {
	user, err := s.adminStore.GetUser(ctx, s.psql.DB(), store.AdminGetUserParams{ID: params.ID})
	if err != nil {
		return GetUserResult{}, fmt.Errorf("get user: %w", err)
	}

	return GetUserResult{
		User: user,
	}, nil
}
