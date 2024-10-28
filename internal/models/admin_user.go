package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type AdminUserRole uint8

const (
	AdminUserRoleInvalid AdminUserRole = iota
	AdminUserRoleEmployee
	AdminUserRoleAdmin
	AdminUserRoleSuperAdmin
)

type AdminUser struct {
	ID           int
	Email        string
	Phone        string
	PasswordHash string
	FirstName    string
	LastName     string
	MiddleName   string
	AvatarURL    string
	Role         AdminUserRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func ConvertAdminUserModelToApi(user AdminUser) api.AdminUser {
	return api.AdminUser{
		ID:         user.ID,
		Email:      api.Email(user.Email),
		Phone:      api.Phone(user.Phone),
		FirstName:  api.Name(user.FirstName),
		LastName:   api.Name(user.LastName),
		MiddleName: api.Name(user.MiddleName),
		AvatarURL: api.OptURL{
			Value: api.URL(user.AvatarURL),
			Set:   user.AvatarURL != "",
		},
		Role:      api.AdminRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
