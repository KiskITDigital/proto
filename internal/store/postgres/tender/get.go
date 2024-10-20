package tender

import (
	"context"
	"database/sql"
	"fmt"

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
