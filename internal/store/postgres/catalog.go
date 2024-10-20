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
