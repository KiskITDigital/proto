package organization

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/deduplicate"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
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
		"o.customer_info",
		"o.contractor_info",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		Offset(params.Offset).
		PlaceholderFormat(sq.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.IsContractor.Set {
		builder = builder.Where(
			sq.Eq{
				"o.is_contractor":       params.IsContractor.Value,
				"o.verification_status": models.VerificationStatusApproved,
				"o.is_banned":           false,
			})
	}

	if params.OrganizationID.Set {
		builder = builder.Where(sq.Eq{"o.id": params.OrganizationID.Value})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	cityIDs := []int{}
	serviceIDs := []int{}
	objectIDs := []int{}

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
			&org.CustomerInfo,
			&org.ContractorInfo,
			&org.CreatedAt,
			&org.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		org.AvatarURL = avatarURL.String

		organizations = append(organizations, org)
		cityIDs = append(cityIDs, org.CustomerInfo.CityIDs...)

		if org.IsContractor {
			cityIDs = append(cityIDs, org.ContractorInfo.CityIDs...)
			serviceIDs = append(serviceIDs, org.ContractorInfo.ServiceIDs...)
			objectIDs = append(objectIDs, org.ContractorInfo.ObjectIDs...)
		}
	}

	cities, err := s.catalogStore.GetCitiesByIDs(ctx, qe, deduplicate.Deduplicate(cityIDs))
	if err != nil {
		return nil, fmt.Errorf("get cities by ids: %w", err)
	}

	objects, err := s.catalogStore.GetObjectsByIDs(ctx, qe, deduplicate.Deduplicate(objectIDs))
	if err != nil {
		return nil, fmt.Errorf("get objects by ids: %w", err)
	}

	services, err := s.catalogStore.GetServicesByIDs(ctx, qe, deduplicate.Deduplicate(serviceIDs))
	if err != nil {
		return nil, fmt.Errorf("get services by ids: %w", err)
	}

	for i, org := range organizations {
		for _, cityID := range org.CustomerInfo.CityIDs {
			if city, ok := cities[cityID]; ok {
				org.CustomerInfo.Cities = append(org.CustomerInfo.Cities, city)
			}
		}

		if org.IsContractor {
			for _, cityID := range org.ContractorInfo.CityIDs {
				if city, ok := cities[cityID]; ok {
					org.ContractorInfo.Cities = append(org.ContractorInfo.Cities, city)
				}
			}

			for _, serviceID := range org.ContractorInfo.ServiceIDs {
				if service, ok := services[serviceID]; ok {
					org.ContractorInfo.Services = append(org.ContractorInfo.Services, service)
				}
			}

			for _, objectID := range org.ContractorInfo.ObjectIDs {
				if object, ok := objects[objectID]; ok {
					org.ContractorInfo.Objects = append(org.ContractorInfo.Objects, object)
				}
			}
		}

		organizations[i] = org
	}

	return organizations, nil
}

func (s *OrganizationStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error) {
	organizations, err := s.Get(ctx, qe, store.OrganizationGetParams{
		OrganizationID: models.NewOptional(id)})
	if err != nil {
		return models.Organization{}, fmt.Errorf("get organizations: %w", err)
	}

	if len(organizations) == 0 {
		return models.Organization{}, errstore.ErrOrganizationNotFound
	}

	return organizations[0], nil
}
