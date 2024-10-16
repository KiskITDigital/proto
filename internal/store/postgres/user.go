package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) Create(ctx context.Context, qe store.QueryExecutor, user models.User) (models.User, error) {
	builder := squirrel.
		Insert("users").
		Columns(
			"organization_id",
			"email",
			"phone",
			"password_hash",
			"totp_salt",
			"first_name",
			"last_name",
			"middle_name",
			"avatar_url",
			"verified",
			"email_verified",
			"role",
			"is_contractor",
		).
		Values(
			user.OrganizationID,
			user.Email,
			user.Phone,
			user.PasswordHash,
			user.TOTPSalt,
			user.FirstName,
			user.LastName,
			user.MiddleName,
			user.AvatarURL,
			user.Verified,
			user.EmailVerified,
			user.Role,
			user.IsContractor,
		).
		Suffix(`
			RETURNING
				id,
				organization_id,
				email,
				password_hash,
				totp_salt,
				first_name,
				last_name,
				middle_name,
				avatar_url,
				verified,
				email_verified,
				role,
				is_contractor
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdUser models.User

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdUser.ID,
		&createdUser.OrganizationID,
		&createdUser.Email,
		&createdUser.Phone,
		&createdUser.PasswordHash,
		&createdUser.TOTPSalt,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.MiddleName,
		&createdUser.AvatarURL,
		&createdUser.Verified,
		&createdUser.EmailVerified,
		&createdUser.Role,
		&createdUser.IsContractor,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("query row: %w", err)
	}

	return createdUser, nil
}
