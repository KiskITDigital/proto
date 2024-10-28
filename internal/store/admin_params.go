package store

import "time"

type AdminGetUserParams struct {
	Email string
	ID    int
}

type AdminCreateUserParams struct {
	Email        string
	Phone        string
	PasswordHash string
	FirstName    string
	LastName     string
	MiddleName   string
	AvatarURL    string
	Role         int
}

type AdminCreateSessionParams struct {
	ID        string
	UserID    int
	IPAddress string
	ExpiresAt time.Time
}
