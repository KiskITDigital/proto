package postgres

import (
	"context"
	"database/sql"
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

func (s *UserStore) Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error) {
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
			params.OrganizationID,
			params.Email,
			params.Phone,
			params.PasswordHash,
			params.TOTPSalt,
			params.FirstName,
			params.LastName,
			params.MiddleName,
			sql.NullString{Valid: params.AvatarURL != "", String: params.AvatarURL},
			params.Verified,
			params.EmailVerified,
			params.Role,
			params.IsContractor,
		).
		Suffix(`
			RETURNING
				id,
				organization_id,
				email,
				phone,
				password_hash,
				totp_salt,
				first_name,
				last_name,
				middle_name,
				avatar_url,
				verified,
				email_verified,
				role,
				is_contractor,
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
		&createdUser.Organization.ID,
		&createdUser.Email,
		&createdUser.Phone,
		&createdUser.PasswordHash,
		&createdUser.TOTPSalt,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.MiddleName,
		&avatarURL,
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

	createdUser.AvatarURL = avatarURL.String

	return createdUser, nil
}

func (s *UserStore) Get(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) (models.User, error) {
	builder := squirrel.
		Select(
			"id",
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
			"created_at",
			"updated_at",
		).
		From("users").
		PlaceholderFormat(squirrel.Dollar)

	if params.Email != "" {
		builder.Where(squirrel.Eq{"email": params.Email})
	}

	if params.ID != 0 {
		builder.Where(squirrel.Eq{"id": params.ID})
	}

	var (
		user      models.User
		avatarURL sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&user.ID,
		&user.Organization.ID,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.TOTPSalt,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&avatarURL,
		&user.Verified,
		&user.EmailVerified,
		&user.Role,
		&user.IsContractor,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("query row: %w", err)
	}

	user.AvatarURL = avatarURL.String

	return user, nil
}

func (s *UserStore) GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) (models.User, error) {
	builder := squirrel.
		Select(
			"u.id",
			"u.organization_id",
			"u.email",
			"u.phone",
			"u.password_hash",
			"u.totp_salt",
			"u.first_name",
			"u.last_name",
			"u.middle_name",
			"u.avatar_url",
			"u.verified",
			"u.email_verified",
			"u.role",
			"u.is_contractor",
			"u.created_at",
			"u.updated_at",
			"o.id",
			"o.brand_name",
			"o.full_name",
			"o.short_name",
			"o.inn",
			"o.okpo",
			"o.ogrn",
			"o.kpp",
			"o.tax_code",
			"o.address",
			"o.avatar_url",
			"o.emails",
			"o.phones",
			"o.messengers",
			"o.created_at",
			"o.updated_at",
		).
		From("users AS u").
		Join("organizations AS o ON u.organization_id = o.id").
		PlaceholderFormat(squirrel.Dollar)

	if params.Email != "" {
		builder.Where(squirrel.Eq{"email": params.Email})
	}

	if params.ID != 0 {
		builder.Where(squirrel.Eq{"id": params.ID})
	}

	var (
		user                  models.User
		userAvatarURL         sql.NullString
		organizationAvatarURL sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&user.ID,
		&user.Organization.ID,
		&user.Email,
		&user.Phone,
		&user.PasswordHash,
		&user.TOTPSalt,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&userAvatarURL,
		&user.Verified,
		&user.EmailVerified,
		&user.Role,
		&user.IsContractor,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Organization.ID,
		&user.Organization.BrandName,
		&user.Organization.FullName,
		&user.Organization.ShortName,
		&user.Organization.INN,
		&user.Organization.OKPO,
		&user.Organization.OGRN,
		&user.Organization.KPP,
		&user.Organization.TaxCode,
		&user.Organization.Address,
		&organizationAvatarURL,
		&user.Organization.Emails,
		&user.Organization.Phones,
		&user.Organization.Messengers,
		&user.Organization.CreatedAt,
		&user.Organization.UpdatedAt,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("query row: %w", err)
	}

	user.AvatarURL = userAvatarURL.String
	user.Organization.AvatarURL = organizationAvatarURL.String

	return user, nil
}
