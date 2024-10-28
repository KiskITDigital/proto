package admin

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Store) GetUser(ctx context.Context, qe store.QueryExecutor, params store.AdminGetUserParams) (models.AdminUser, error) {
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

	if params.Email != "" {
		builder = builder.Where(squirrel.Eq{"email": params.Email})
	}

	if params.ID != 0 {
		builder = builder.Where(squirrel.Eq{"id": params.ID})
	}

	var (
		adminUser models.AdminUser
		avatarURL sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&adminUser.ID,
		&adminUser.Email,
		&adminUser.Phone,
		&adminUser.PasswordHash,
		&adminUser.FirstName,
		&adminUser.LastName,
		&adminUser.MiddleName,
		&avatarURL,
		&adminUser.Role,
		&adminUser.CreatedAt,
		&adminUser.UpdatedAt,
	)
	if err != nil {
		return models.AdminUser{}, fmt.Errorf("query row: %w", err)
	}

	adminUser.AvatarURL = avatarURL.String

	return adminUser, nil
}
