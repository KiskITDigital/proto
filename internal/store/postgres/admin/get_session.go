package admin

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Store) GetSession(ctx context.Context, qe store.QueryExecutor, sessionID string) (models.Session, error) {
	builder := squirrel.
		Select(
			"id",
			"user_id",
			"ip_address",
			"created_at",
			"expires_at",
		).
		From("admin.sessions").
		Where(squirrel.Eq{"id": sessionID}).
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
