package catalog

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

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
