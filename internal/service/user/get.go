package user

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, userID int) (models.User, error) {
	users, err := s.userStore.GetWithOrganiztion(ctx, s.psql.DB(), store.UserGetParams{ID: userID})
	if err != nil {
		return models.User{}, err
	}
	return users[0], nil
}

func (s *Service) Get(ctx context.Context) ([]models.User, error) {
	return s.userStore.GetWithOrganiztion(ctx, s.psql.DB(), store.UserGetParams{})
}
