package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type UserRole uint8

const (
	UserRoleInvalid = iota
	UserRoleUser
	UserRoleEmployee
	UserRoleAdmin
	UserRoleSuperAdmin
)

func (r UserRole) ToApi() string {
	switch r {
	case UserRoleUser:
		return "user"
	case UserRoleEmployee:
		return "employee"
	case UserRoleAdmin:
		return "admin"
	case UserRoleSuperAdmin:
		return "super_admin"
	default:
		return "invalid"
	}
}

type User struct {
	ID            int
	Email         string
	Phone         string
	PasswordHash  string
	TOTPSalt      string
	FirstName     string
	LastName      string
	MiddleName    string
	AvatarURL     string
	EmailVerified bool
	Role          UserRole
	IsBanned      bool
	Organization  Organization
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func ConvertUserModelToApi(user User) api.User {
	return api.User{
		ID:            user.ID,
		Email:         api.Email(user.Email),
		Phone:         api.Phone(user.Phone),
		FirstName:     api.Name(user.FirstName),
		LastName:      api.Name(user.LastName),
		MiddleName:    api.Name(user.MiddleName),
		AvatarURL:     api.OptURL{Value: api.URL(user.AvatarURL), Set: user.AvatarURL != ""},
		EmailVerified: user.EmailVerified,
		Role:          api.Role(user.Role.ToApi()),
		Organization: api.OptOrganization{
			Value: api.Organization(ConvertOrganizationModelToApi(user.Organization)),
			Set:   user.Organization.ID != 0,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
