package user

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, userID int) (models.RegularUser, error) {
	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{ID: userID})
	if err != nil {
		return models.RegularUser{}, err
	}

	if len(users) == 0 {
		return models.RegularUser{}, cerr.Wrap(
			fmt.Errorf("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %d id not found", userID),
			nil,
		)
	}

	return models.RegularUser{User: users[0].User, Organization: users[0].Organization}, nil
}

func (s *Service) Get(ctx context.Context) ([]models.FullUser, error) {
	return s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{})
}
