package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type TenderStore struct {
}

func NewTenderStore() *TenderStore {
	return &TenderStore{}
}

func (s *TenderStore) Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (models.Tender, error) {
	builder := squirrel.
		Insert("tenders").
		Columns(
			"name",
			"price",
			"is_contract_price",
			"is_nds_price",
			"city_id",
			"floor_space",
			"description",
			"wishes",
			"specification",
			"attachments",
			"reception_start",
			"reception_end",
			"work_start",
			"work_end",
			"organization_id",
		).
		Values(
			params.Name,
			params.Price,
			params.IsContractPrice,
			params.IsNDSPrice,
			params.CityID,
			params.FloorSpace,
			params.Description,
			params.Wishes,
			params.Specification,
			pq.Array(params.Attachments),
			params.ReceptionStart,
			params.ReceptionEnd,
			params.WorkStart,
			params.WorkEnd,
			params.OrganizationID,
		).
		Suffix(`
			RETURNING
				id,
				name,
				price,
				is_contract_price,
				is_nds_price,
				city_id,
				floor_space,
				description,
				wishes,
				specification,
				attachments,
				verified,
				reception_start,
				reception_end,
				work_start,
				work_end,
				organization_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdTender models.Tender

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdTender.ID,
		&createdTender.Name,
		&createdTender.Price,
		&createdTender.IsContractPrice,
		&createdTender.IsNDSPrice,
		&createdTender.City.ID,
		&createdTender.FloorSpace,
		&createdTender.Description,
		&createdTender.Wishes,
		&createdTender.Specification,
		pq.Array(&createdTender.Attachments),
		&createdTender.Verified,
		&createdTender.ReceptionStart,
		&createdTender.ReceptionEnd,
		&createdTender.WorkStart,
		&createdTender.WorkEnd,
		&createdTender.Organization.ID,
	)
	if err != nil {
		return models.Tender{}, fmt.Errorf("query row: %w", err)
	}

	return createdTender, nil
}

func (s *TenderStore) AppendTenderServies(ctx context.Context, qe store.QueryExecutor, params store.TenderServicesCreateParams) error {
	builder := squirrel.
		Insert("tender_services").
		Columns(
			"tender_id",
			"service_id",
		).
		PlaceholderFormat(squirrel.Dollar)

	for _, serviceID := range params.ServicesIDs {
		builder = builder.Values(params.TenderID, serviceID)
	}

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (s *TenderStore) AppendTenderObjects(ctx context.Context, qe store.QueryExecutor, params store.TenderObjectsCreateParams) error {
	builder := squirrel.
		Insert("tender_objects").
		Columns(
			"tender_id",
			"object_id",
		).
		PlaceholderFormat(squirrel.Dollar)

	for _, objectID := range params.ObjectsIDs {
		builder = builder.Values(params.TenderID, objectID)
	}

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}

func (s *TenderStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error) {
	builder := squirrel.
		Select(
			"t.id",
			"t.name",
			"t.price",
			"t.is_contract_price",
			"t.is_nds_price",
			"c.name",
			"r.name",
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
		AvatarURL     sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdTender.ID,
		&createdTender.Name,
		&createdTender.Price,
		&createdTender.IsContractPrice,
		&createdTender.IsNDSPrice,
		&createdTender.City.Name,
		&createdTender.Region.Name,
		&createdTender.FloorSpace,
		&createdTender.Description,
		&createdTender.Wishes,
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

	services, err := s.GetTenderServices(ctx, qe, createdTender.ID)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get services: %w", err)
	}

	createdTender.Services = services

	objects, err := s.GetTenderObjects(ctx, qe, createdTender.ID)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get objects: %w", err)
	}

	createdTender.Objects = objects

	return createdTender, nil
}

func (s *TenderStore) GetTenderServices(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.TenderService, error) {
	query := `WITH RECURSIVE service_hierarchy AS (
		SELECT 
			s.id AS service_id,
			s.name AS service_name,
			s.parent_id
		FROM 
			services s
		JOIN 
			tender_services ts ON s.id = ts.service_id
		WHERE 
			ts.tender_id = $1

		UNION ALL

		SELECT 
			s.id AS service_id,
			s.name AS service_name,
			s.parent_id
		FROM 
			services s
		JOIN 
			service_hierarchy sh ON s.id = sh.parent_id
	)

	SELECT 
		service_id,
		service_name,
		parent_id
	FROM 
		service_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, tenderID)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	services := []models.TenderService{}

	for rows.Next() {
		var (
			service  models.TenderService
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&service.ID,
			&service.Name,
			&parentID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		service.ParentID = int(parentID.Int64)

		services = append(services, service)
	}

	return services, nil
}

func (s *TenderStore) GetTenderObjects(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.TenderObject, error) {
	query := `WITH RECURSIVE object_hierarchy AS (
		SELECT 
			o.id AS object_id,
			o.name AS object_name,
			o.parent_id
		FROM 
			objects o
		JOIN 
			tender_objects tobj ON o.id = tobj.object_id
		WHERE 
			tobj.tender_id = $1

		UNION ALL

		SELECT 
			o.id AS object_id,
			o.name AS object_name,
			o.parent_id
		FROM 
			objects o
		JOIN 
			object_hierarchy sh ON o.id = sh.parent_id
	)

	SELECT 
		object_id,
		object_name,
		parent_id
	FROM 
		object_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, tenderID)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	objects := []models.TenderObject{}

	for rows.Next() {
		var (
			object   models.TenderObject
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&object.ID,
			&object.Name,
			&parentID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		object.ParentID = int(parentID.Int64)

		objects = append(objects, object)
	}

	return objects, nil
}
