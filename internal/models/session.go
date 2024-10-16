package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	RefreshToken uuid.UUID
	UserID       int
	IPAddress    string
	CreatedAt    time.Time
	ExpiresAt    time.Time
}
