package catalog

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

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

func (s *CatalogStore) ListCities(ctx context.Context, qe store.QueryExecutor, name string) ([]models.City, error) {
	builder := squirrel.
		Select(
			"c.id",
			"c.name",
			"c.region_id",
			"r.name",
		).
		From("cities c").
		Join("regions r ON r.id = c.region_id").
		PlaceholderFormat(squirrel.Dollar)

	if name != "" {
		builder = builder.Where(squirrel.ILike{"c.name": fmt.Sprintf("%s%%", name)})
	}

	var cities []models.City

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	for rows.Next() {
		var city models.City

		err = rows.Scan(
			&city.ID,
			&city.Name,
			&city.Region.ID,
			&city.Region.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		cities = append(cities, city)
	}

	return cities, nil
}
