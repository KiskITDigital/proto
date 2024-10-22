package user

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, tenderID int) (models.User, error) {
	return s.userStore.GetWithOrganiztion(ctx, s.psql.DB(), store.UserGetParams{ID: tenderID})
}
