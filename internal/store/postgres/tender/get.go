package tender

import (
	"context"
	"database/sql"
	"fmt"
	"maps"
	"slices"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error) {
	builder := squirrel.
		Select(
			"t.id",
			"t.name",
			"t.price",
			"t.is_contract_price",
			"t.is_nds_price",
			"t.is_draft",
			"c.name",
			"c.id",
			"r.name",
			"r.id",
			"t.floor_space",
			"t.description",
			"t.wishes",
			"t.specification",
			"t.attachments",
			"t.verified",
			"t.reception_start",
			"t.reception_end",
			"t.work_start",
			"t.work_end",
			"t.created_at",
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
		From("tenders AS t").
		Join("cities AS c ON c.id = t.city_id").
		Join("regions AS r ON r.id = c.region_id").
		Join("organizations AS o ON o.id = t.organization_id").
		Where(squirrel.Eq{"t.id": id}).
		PlaceholderFormat(squirrel.Dollar)

	var (
		createdTender models.Tender
		description   sql.NullString
		wishes        sql.NullString
		AvatarURL     sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdTender.ID,
		&createdTender.Name,
		&createdTender.Price,
		&createdTender.IsContractPrice,
		&createdTender.IsNDSPrice,
		&createdTender.IsDraft,
		&createdTender.City.Name,
		&createdTender.City.ID,
		&createdTender.City.Region.Name,
		&createdTender.City.Region.ID,
		&createdTender.FloorSpace,
		&description,
		&wishes,
		&createdTender.Specification,
		pq.Array(&createdTender.Attachments),
		&createdTender.Verified,
		&createdTender.ReceptionStart,
		&createdTender.ReceptionEnd,
		&createdTender.WorkStart,
		&createdTender.WorkEnd,
		&createdTender.CreatedAt,
		&createdTender.Organization.ID,
		&createdTender.Organization.BrandName,
		&createdTender.Organization.FullName,
		&createdTender.Organization.ShortName,
		&createdTender.Organization.INN,
		&createdTender.Organization.OKPO,
		&createdTender.Organization.OGRN,
		&createdTender.Organization.KPP,
		&createdTender.Organization.TaxCode,
		&createdTender.Organization.Address,
		&AvatarURL,
		&createdTender.Organization.Emails,
		&createdTender.Organization.Phones,
		&createdTender.Organization.Messengers,
		&createdTender.Organization.CreatedAt,
		&createdTender.Organization.UpdatedAt,
	)
	if err != nil {
		return models.Tender{}, fmt.Errorf("query row: %w", err)
	}

	createdTender.Organization.AvatarURL = AvatarURL.String
	createdTender.Description = description.String
	createdTender.Wishes = wishes.String

	services, err := s.GetTendersServices(ctx, qe, []int{createdTender.ID})
	if err != nil {
		return models.Tender{}, fmt.Errorf("get services: %w", err)
	}

	createdTender.Services = services

	objects, err := s.GetTendersObjects(ctx, qe, []int{createdTender.ID})
	if err != nil {
		return models.Tender{}, fmt.Errorf("get objects: %w", err)
	}

	createdTender.Objects = objects

	return createdTender, nil
}

func (s *TenderStore) List(ctx context.Context, qe store.QueryExecutor, params store.TenderGetParams) ([]models.Tender, error) {
	builder := squirrel.
		Select(
			"t.id",
			"t.organization_id",
			"t.winner_organization_id",
			"t.city_id",
			"t.services_ids",
			"t.objects_ids",
			"t.name",
			"t.price",
			"t.is_contract_price",
			"t.is_nds_price",
			"t.floor_space",
			"t.description",
			"t.wishes",
			"t.specification",
			"t.attachments",
			"t.status",
			"t.verification_status",
			"t.is_draft",
			"t.reception_start",
			"t.reception_end",
			"t.work_start",
			"t.work_end",
			"t.created_at",
			"t.updated_at",
			"c.name",
			"c.id",
			"r.name",
			"r.id",
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
		).
		From("tenders AS t").
		Join("cities AS c ON c.id = t.city_id").
		Join("regions AS r ON r.id = c.region_id").
		Join("organizations AS o ON o.id = t.organization_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.OrganizationID.Set {
		builder = builder.Where(squirrel.Eq{"t.organization_id": params.OrganizationID.Value})
	}

	if !params.WithDrafts {
		builder = builder.Where(squirrel.Eq{"t.is_draft": false})
	}

	if params.VerifiedOnly {
		builder = builder.Where(squirrel.Eq{"t.verification_status": models.VerificationStatusApproved})
	}

	tenders := make(map[int]models.Tender)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			tender      models.Tender
			serviceIDs  []int
			objectsIDs  []int
			description sql.NullString
			wishes      sql.NullString
			AvatarURL   sql.NullString
		)

		err = rows.Scan(
			&tender.ID,
			&tender.Organization.ID,
			&tender.City.ID,
			pq.Array(&serviceIDs),
			pq.Array(&objectsIDs),
			&tender.Name,
			&tender.Price,
			&tender.IsContractPrice,
			&tender.IsNDSPrice,
			&tender.FloorSpace,
			&description,
			&wishes,
			&tender.Specification,
			pq.Array(&tender.Attachments),
			&tender.Status,
			&tender.VerificationStatus,
			&tender.IsDraft,
			&tender.ReceptionStart,
			&tender.ReceptionEnd,
			&tender.WorkStart,
			&tender.WorkEnd,
			&tender.CreatedAt,
			&tender.UpdatedAt,
			&tender.City.Name,
			&tender.City.ID,
			&tender.City.Region.Name,
			&tender.City.Region.ID,
			&tender.Organization.ID,
			&tender.Organization.BrandName,
			&tender.Organization.FullName,
			&tender.Organization.ShortName,
			&tender.Organization.INN,
			&tender.Organization.OKPO,
			&tender.Organization.OGRN,
			&tender.Organization.KPP,
			&tender.Organization.TaxCode,
			&tender.Organization.Address,
			&AvatarURL,
			pq.Array(&tender.Organization.Emails),
			pq.Array(&tender.Organization.Phones),
			pq.Array(&tender.Organization.Messengers),
			&tender.Organization.VerificationStatus,
			&tender.Organization.IsContractor,
			&tender.Organization.IsBanned,
			&tender.Organization.CustomerInfo,
			&tender.Organization.ContractorInfo,
			&tender.Organization.CreatedAt,
			&tender.Organization.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		tender.Organization.AvatarURL = AvatarURL.String
		tender.Description = description.String
		tender.Wishes = wishes.String

		tenders[tender.ID] = tender
	}

	tendersIDs := slices.Collect(maps.Keys(tenders))

	services, err := s.GetTendersServices(ctx, qe, tendersIDs)
	if err != nil {
		return nil, fmt.Errorf("get services: %w", err)
	}

	for _, service := range services {
		tender := tenders[service.TenderID]
		tender.Services = append(tender.Services, service)
		tenders[service.TenderID] = tender
	}

	objects, err := s.GetTendersObjects(ctx, qe, tendersIDs)
	if err != nil {
		return nil, fmt.Errorf("get objects: %w", err)
	}

	for _, object := range objects {
		tender := tenders[object.TenderID]
		tender.Objects = append(tender.Objects, object)
		tenders[object.TenderID] = tender
	}

	return slices.Collect(maps.Values(tenders)), nil
}
