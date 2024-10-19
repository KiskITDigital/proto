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

func (s *SessionStore) Create(ctx context.Context, qe store.QueryExecutor, params store.SessionCreateParams) (models.Session, error) {
	builder := squirrel.
		Insert("sessions").
		Columns(
			"id",
			"user_id",
			"ip_address",
			"expires_at",
		).
		Values(
			params.ID,
			params.UserID,
			params.IPAddress,
			params.ExpiresAt,
		).
		Suffix(`
			RETURNING
				id,
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

func (s *SessionStore) Get(ctx context.Context, qe store.QueryExecutor, params store.SessionGetParams) (models.Session, error) {
	builder := squirrel.
		Select(
			"id",
			"user_id",
			"ip_address",
			"created_at",
			"expires_at",
		).
		From("sessions").
		Where(squirrel.Eq{"id": params.ID}).
		PlaceholderFormat(squirrel.Dollar)

	var session models.Session

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&session.ID,
		&session.UserID,
		&session.IPAddress,
		&session.CreatedAt,
		&session.ExpiresAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("query row: %w", err)
	}

	return session, nil
}

func (s *SessionStore) Update(ctx context.Context, qe store.QueryExecutor, params store.SessionUpdateParams) (models.Session, error) {
	builder := squirrel.
		Update("sessions").
		Set("expires_at", params.ExpiresAt).
		Where(squirrel.Eq{"id": params.ID}).
		Suffix(`
			RETURNING
				id,
				user_id,
				ip_address,
				created_at,
				expires_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var session models.Session

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&session.ID,
		&session.UserID,
		&session.IPAddress,
		&session.CreatedAt,
		&session.ExpiresAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("query row: %w", err)
	}

	return session, nil
}
