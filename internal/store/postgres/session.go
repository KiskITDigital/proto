package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type SessionStore struct {
}

func NewSessionStore() *SessionStore {
	return &SessionStore{}
}

func (s *SessionStore) Create(ctx context.Context, qe store.QueryExecutor, session models.Session) (models.Session, error) {
	builder := squirrel.
		Insert("sessions").
		Columns(
			"id",
			"user_id",
			"ip_address",
			"expires_at",
		).
		Values(
			session.ID,
			session.UserID,
			session.IPAddress,
			session.ExpiresAt,
		).
		Suffix(`
			RETURNING
				refresh_token,
				user_id,
				ip_address,
				created_at,
				expires_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdSession models.Session

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdSession.ID,
		&createdSession.UserID,
		&createdSession.IPAddress,
		&createdSession.CreatedAt,
		&createdSession.ExpiresAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("query row: %w", err)
	}

	return createdSession, nil
}
