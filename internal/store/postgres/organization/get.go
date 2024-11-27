package organization

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) Get(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetParams) ([]models.Organization, error) {
	builder := sq.Select(
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
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(sq.Dollar)

	if params.IsContractor.Set {
		builder = builder.Where(
			sq.Eq{
				"o.is_contractor":       params.IsContractor.Value,
				"o.verification_status": models.VerificationStatusApproved,
				"o.is_banned":           false,
			})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	organizations := []models.Organization{}
	for rows.Next() {
		var (
			org       models.Organization
			avatarURL sql.NullString
		)

		if err = rows.Scan(
			&org.ID,
			&org.BrandName,
			&org.FullName,
			&org.ShortName,
			&org.INN,
			&org.OKPO,
			&org.OGRN,
			&org.KPP,
			&org.TaxCode,
			&org.Address,
			&avatarURL,
			&org.Emails,
			&org.Phones,
			&org.Messengers,
			&org.VerificationStatus,
			&org.IsContractor,
			&org.IsBanned,
			&org.CreatedAt,
			&org.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		org.AvatarURL = avatarURL.String

		organizations = append(organizations, org)
	}

	return organizations, nil
}

func (s *OrganizationStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error) {
	builder := sq.Select(
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
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.created_at",
		"o.updated_at").
		From("organizations AS o").
		Where(sq.Eq{"o.id": id}).
		PlaceholderFormat(sq.Dollar)

	var (
		org       models.Organization
		avatarURL sql.NullString
	)

	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&org.ID,
		&org.BrandName,
		&org.FullName,
		&org.ShortName,
		&org.INN,
		&org.OKPO,
		&org.OGRN,
		&org.KPP,
		&org.TaxCode,
		&org.Address,
		&avatarURL,
		&org.Emails,
		&org.Phones,
		&org.Messengers,
		&org.VerificationStatus,
		&org.IsContractor,
		&org.IsBanned,
		&org.CreatedAt,
		&org.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Organization{}, errors.New("organization not found")
		}
		return models.Organization{}, fmt.Errorf("query row: %w", err)
	}

	org.AvatarURL = avatarURL.String

	return org, nil
}
