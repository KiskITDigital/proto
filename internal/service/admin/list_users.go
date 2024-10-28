package admin

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type ListUsersResult struct {
	Users []models.AdminUser
}

func (s *Service) ListUsers(ctx context.Context) (ListUsersResult, error) {
	users, err := s.adminStore.ListUsers(ctx, s.psql.DB())
	if err != nil {
		return ListUsersResult{}, fmt.Errorf("list users: %w", err)
	}

	return ListUsersResult{
		Users: users,
	}, nil
}
