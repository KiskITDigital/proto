package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type UserRole uint8

const (
	UserRoleInvalid = iota
	UserRoleUser
	UserRoleAdmin
)

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
	Verified      bool
	EmailVerified bool
	Role          UserRole
	IsContractor  bool
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
		AvatarURL:     api.URL(user.AvatarURL),
		Verified:      user.Verified,
		EmailVerified: user.EmailVerified,
		Role:          api.Role(user.Role),
		IsContractor:  user.IsContractor,
		IsBanned:      user.IsBanned,
		Organization: api.OptOrganization{
			Value: api.Organization(ConvertOrganizationModelToApi(user.Organization)),
			Set:   user.Organization.ID != 0,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
