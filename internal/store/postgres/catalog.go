package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type CatalogStore struct {
}

func NewCatalogStore() *CatalogStore {
	return &CatalogStore{}
}

// select s.id, s.name, s.parent_id
// from services AS s
// left join services s2 ON s2.id = s.parent_id;

func (s *CatalogStore) GetServices(ctx context.Context, qe store.QueryExecutor) (models.CatalogServices, error) {
	builder := squirrel.
		Select(
			"s.id",
			"s.name",
			"s.parent_id",
		).
		From("services s").
		LeftJoin("services s2 ON s2.id = s.parent_id;").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var services models.CatalogServices

	for rows.Next() {
		var (
			service  models.CatalogService
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

func (s *CatalogStore) GetObjects(ctx context.Context, qe store.QueryExecutor) (models.CatalogObjects, error) {
	builder := squirrel.
		Select(
			"o.id",
			"o.name",
			"o.parent_id",
		).
		From("objects o").
		LeftJoin("objects o2 ON o2.id = o.parent_id;").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var objects models.CatalogObjects

	for rows.Next() {
		var (
			object   models.CatalogObject
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

func (s *CatalogStore) CreateRegion(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateRegionParams) (models.Region, error) {
	builder := squirrel.
		Insert("regions").
		Columns(
			"name",
		).
		Values(
			params.Name,
		).
		Suffix(`
			RETURNING
				id,
				name
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdRegion models.Region

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdRegion.ID,
		&createdRegion.Name,
	)
	if err != nil {
		return models.Region{}, fmt.Errorf("query row: %w", err)
	}

	return createdRegion, nil
}

func (s *CatalogStore) CreateCity(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateCityParams) (models.City, error) {
	builder := squirrel.
		Insert("cities").
		Columns(
			"name",
			"region_id",
		).
		Values(
			params.Name,
			params.RegionID,
		).
		Suffix(`
			RETURNING
				id,
				name,
				region_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdCity models.City

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdCity.ID,
		&createdCity.Name,
		&createdCity.Region.ID,
	)
	if err != nil {
		return models.City{}, fmt.Errorf("query row: %w", err)
	}

	return createdCity, nil
}

func (s *CatalogStore) GetRegionByID(ctx context.Context, qe store.QueryExecutor, regionID int) (models.Region, error) {
	builder := squirrel.
		Select(
			"id",
			"name",
		).
		From("regions").
		Where(squirrel.Eq{"id": regionID}).
		PlaceholderFormat(squirrel.Dollar)

	var region models.Region

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&region.ID,
		&region.Name,
	)
	if err != nil {
		return models.Region{}, fmt.Errorf("query row: %w", err)
	}

	return region, nil
}

func (s *CatalogStore) GetCityByID(ctx context.Context, qe store.QueryExecutor, cityID int) (models.City, error) {
	builder := squirrel.
		Select(
			"c.id",
			"c.name",
			"c.region_id",
			"r.name",
		).
		From("cities c").
		Join("regions r ON r.id = c.region_id").
		Where(squirrel.Eq{"c.id": cityID}).
		PlaceholderFormat(squirrel.Dollar)

	var city models.City

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&city.ID,
		&city.Name,
		&city.Region.ID,
		&city.Region.Name,
	)
	if err != nil {
		return models.City{}, fmt.Errorf("query row: %w", err)
	}

	return city, nil
}

func (s *CatalogStore) CreateObject(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateObjectParams) (models.CatalogObject, error) {
	builder := squirrel.
		Insert("objects").
		Columns(
			"name",
			"parent_id",
		).
		Values(
			params.Name,
			sql.NullInt64{Int64: int64(params.ParentID), Valid: params.ParentID != 0},
		).
		Suffix(`
			RETURNING
				id,
				name,
				parent_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		object   models.CatalogObject
		parentID sql.NullInt64
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&object.ID,
		&object.Name,
		&parentID,
	)
	if err != nil {
		return models.CatalogObject{}, fmt.Errorf("query row: %w", err)
	}

	object.ParentID = int(parentID.Int64)

	return object, nil
}

func (s *CatalogStore) CreateService(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateServiceParams) (models.CatalogService, error) {
	builder := squirrel.
		Insert("services").
		Columns(
			"name",
			"parent_id",
		).
		Values(
			params.Name,
			sql.NullInt64{Int64: int64(params.ParentID), Valid: params.ParentID != 0},
		).
		Suffix(`
			RETURNING
				id,
				name,
				parent_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		service  models.CatalogService
		parentID sql.NullInt64
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&service.ID,
		&service.Name,
		&parentID,
	)
	if err != nil {
		return models.CatalogService{}, fmt.Errorf("query row: %w", err)
	}

	service.ParentID = int(parentID.Int64)

	return service, nil
}
