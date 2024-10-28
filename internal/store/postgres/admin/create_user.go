package admin

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Store) CreateUser(ctx context.Context, qe store.QueryExecutor, params store.AdminCreateUserParams) (models.AdminUser, error) {
	builder := squirrel.
		Insert("admin.users").
		Columns(
			"email",
			"phone",
			"password_hash",
			"first_name",
			"last_name",
			"middle_name",
			"avatar_url",
			"role",
		).
		Values(
			params.Email,
			params.Phone,
			params.PasswordHash,
			params.FirstName,
			params.LastName,
			params.MiddleName,
			sql.NullString{Valid: params.AvatarURL != "", String: params.AvatarURL},
			params.Role,
		).
		Suffix(`
			RETURNING
				id,
				email,
				phone,
				password_hash,
				first_name,
				last_name,
				middle_name,
				avatar_url,
				role,
				created_at,
				updated_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		createdUser models.AdminUser
		avatarURL   sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdUser.ID,
		&createdUser.Email,
		&createdUser.Phone,
		&createdUser.PasswordHash,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.MiddleName,
		&avatarURL,
		&createdUser.Role,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return models.AdminUser{}, fmt.Errorf("query row: %w", err)
	}

	createdUser.AvatarURL = avatarURL.String

	return createdUser, nil
}
