package models

import "time"

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
