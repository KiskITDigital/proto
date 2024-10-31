package organization

import (
	"context"
	"database/sql"
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
		"o.verified",
		"o.is_contractor",
		"o.is_banned",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		PlaceholderFormat(sq.Dollar)

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
			&org.Verified,
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
