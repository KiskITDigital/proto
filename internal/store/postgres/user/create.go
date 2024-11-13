package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *UserStore) Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error) {
	builder := squirrel.
		Insert("users").
		Columns(
			"email",
			"phone",
			"password_hash",
			"totp_salt",
			"first_name",
			"last_name",
			"middle_name",
			"avatar_url",
			"email_verified",
			"role",
			"is_banned",
		).
		Values(
			params.Email,
			params.Phone,
			params.PasswordHash,
			params.TOTPSalt,
			params.FirstName,
			params.LastName,
			params.MiddleName,
			sql.NullString{Valid: params.AvatarURL != "", String: params.AvatarURL},
			false,
			params.Role,
			false,
		).
		Suffix(`
			RETURNING
				id,
				email,
				phone,
				password_hash,
				totp_salt,
				first_name,
				last_name,
				middle_name,
				avatar_url,
				email_verified,
				role,
				created_at,
				updated_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		createdUser models.User
		avatarURL   sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdUser.ID,
		&createdUser.Email,
		&createdUser.Phone,
		&createdUser.PasswordHash,
		&createdUser.TOTPSalt,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.MiddleName,
		&avatarURL,
		&createdUser.EmailVerified,
		&createdUser.Role,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("query row: %w", err)
	}

	createdUser.AvatarURL = avatarURL.String

	return createdUser, nil
}
