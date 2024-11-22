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

func APIRoleToModel(r api.Role) UserRole {
	switch r {
	case api.RoleUser:
		return UserRoleUser
	case api.RoleEmployee:
		return UserRoleEmployee
	case api.RoleAdmin:
		return UserRoleAdmin
	case api.RoleSuperAdmin:
		return UserRoleSuperAdmin
	default:
		return UserRoleInvalid
	}
}

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
	IsBanned      bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type RegularUser struct {
	User
	Organization Organization
}

type EmployeeUser struct {
	User
	Role     UserRole
	Position string
}

type FullUser struct {
	User
	RegularUser
	EmployeeUser
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
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

func ConvertRegularUserModelToApi(user RegularUser) api.RegularUser {
	return api.RegularUser{
		ID:            user.ID,
		Email:         api.Email(user.Email),
		Phone:         api.Phone(user.Phone),
		FirstName:     api.Name(user.FirstName),
		LastName:      api.Name(user.LastName),
		MiddleName:    api.Name(user.MiddleName),
		AvatarURL:     api.OptURL{Value: api.URL(user.AvatarURL), Set: user.AvatarURL != ""},
		EmailVerified: user.EmailVerified,
		Organization:  ConvertOrganizationModelToApi(user.Organization),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

func ConvertEmployeeUserModelToApi(user EmployeeUser) api.EmployeeUser {
	return api.EmployeeUser{
		ID:            user.ID,
		Email:         api.Email(user.Email),
		Phone:         api.Phone(user.Phone),
		FirstName:     api.Name(user.FirstName),
		LastName:      api.Name(user.LastName),
		MiddleName:    api.Name(user.MiddleName),
		AvatarURL:     api.OptURL{Value: api.URL(user.AvatarURL), Set: user.AvatarURL != ""},
		EmailVerified: user.EmailVerified,
		Role:          api.Role(user.Role.ToApi()),
		Position:      user.Position,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}
