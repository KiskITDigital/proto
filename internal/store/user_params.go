package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type UserCreateParams struct {
	Email         string
	Phone         string
	PasswordHash  string
	TOTPSalt      string
	FirstName     string
	LastName      string
	MiddleName    models.Optional[string]
	EmailVerified bool
	AvatarURL     string
}

type UserCreateEmployeeParams struct {
	UserID    int
	Role      models.UserRole
	Postition string
}

type UserGetParams struct {
	Email string
	ID    int
}

type ResetPasswordParams struct {
	UserID       int
	PasswordHash string
}
