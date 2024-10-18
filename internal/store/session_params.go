package store

import "time"

type SessionCreateParams struct {
	ID        string
	UserID    int
	IPAddress string
	ExpiresAt time.Time
}
