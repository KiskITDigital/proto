package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type OrganizationStore struct {
}

func NewOrganizationStore() *OrganizationStore {
	return &OrganizationStore{}
}

func (s *OrganizationStore) Create(ctx context.Context, qe store.QueryExecutor, organization models.Organization) (models.Organization, error) {
	builder := squirrel.
		Insert("organizations").
		Columns(
			"brand_name",
			"full_name",
			"short_name",
			"inn",
			"okpo",
			"ogrn",
			"kpp",
			"tax_code",
			"address",
			"avatar_url",
			"emails",
			"phones",
			"messengers",
		).
		Values(
			organization.BrandName,
			organization.FullName,
			organization.ShortName,
			organization.INN,
			organization.OKPO,
			organization.ORGN,
			organization.KPP,
			organization.TaxCode,
			organization.Address,
			organization.AvatarURL,
			organization.Emails,
			organization.Phones,
			organization.Messangers,
		).
		Suffix(`
			RETURNING
				id,
				brand_name,
				full_name,
				short_name,
				inn,
				okpo,
				ogrn,
				kpp,
				tax_code,
				address,
				avatar_url,
				emails,
				phones,
				messengers,
				created_at,
				updated_at
	`).
		PlaceholderFormat(squirrel.Dollar)

	var createdOrganization models.Organization

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdOrganization.ID,
		&createdOrganization.BrandName,
		&createdOrganization.FullName,
		&createdOrganization.ShortName,
		&createdOrganization.INN,
		&createdOrganization.OKPO,
		&createdOrganization.ORGN,
		&createdOrganization.KPP,
		&createdOrganization.TaxCode,
		&createdOrganization.Address,
		&createdOrganization.AvatarURL,
		&createdOrganization.Emails,
		&createdOrganization.Phones,
		&createdOrganization.Messangers,
		&createdOrganization.CreatedAt,
		&createdOrganization.UpdatedAt,
	)
	if err != nil {
		return models.Organization{}, fmt.Errorf("query row: %w", err)
	}

	return createdOrganization, nil
}
