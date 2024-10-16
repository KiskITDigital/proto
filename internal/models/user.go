package models

import "time"

type UserRole uint8

const (
	UserRoleInvalid = iota
	UserRoleBanned
	UserRoleUser
	UserRoleAdmin
)

type User struct {
	ID             int
	OrganizationID int
	Email          string
	Phone          string
	PasswordHash   string
	TOTPSalt       string
	FirstName      string
	LastName       string
	MiddleName     string
	AvatarURL      string
	Verified       bool
	EmailVerified  bool
	Role           UserRole
	IsContractor   bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
