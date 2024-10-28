package admin

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Store) ListUsers(ctx context.Context, qe store.QueryExecutor) ([]models.AdminUser, error) {
	builder := squirrel.
		Select(
			"id",
			"email",
			"phone",
			"password_hash",
			"first_name",
			"last_name",
			"middle_name",
			"avatar_url",
			"role",
			"created_at",
			"updated_at",
		).
		From("admin.users").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var users []models.AdminUser

	for rows.Next() {
		var (
			user      models.AdminUser
			avatarURL sql.NullString
		)

		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.PasswordHash,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&avatarURL,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		user.AvatarURL = avatarURL.String

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return users, nil
}
