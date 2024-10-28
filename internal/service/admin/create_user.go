package admin

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type CreateUserParams struct {
	Email      string
	Phone      string
	Password   string
	FirstName  string
	LastName   string
	MiddleName string
	AvatarURL  string
	Role       models.UserRole
}

type CreateUserResult struct {
	User models.AdminUser
}

func (s *Service) CreateUser(ctx context.Context, params CreateUserParams) (CreateUserResult, error) {
	hashedPassword, err := crypto.Password(params.Password)
	if err != nil {
		return CreateUserResult{}, fmt.Errorf("hash password: %w", err)
	}

	user, err := s.adminStore.CreateUser(ctx, s.psql.DB(), store.AdminCreateUserParams{
		Email:        params.Email,
		Phone:        params.Phone,
		PasswordHash: hashedPassword,
		FirstName:    params.FirstName,
		LastName:     params.LastName,
		MiddleName:   params.MiddleName,
		AvatarURL:    params.AvatarURL,
		Role:         int(params.Role),
	})
	if err != nil {
		return CreateUserResult{}, fmt.Errorf("create user: %w", err)
	}

	return CreateUserResult{
		User: user,
	}, nil
}
