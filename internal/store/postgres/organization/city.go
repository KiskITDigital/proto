package organization

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) GetCities(ctx context.Context, qe store.QueryExecutor, cityIDs []int) (map[int]models.City, error) {
	builder := sq.Select(
		"c.id",
		"c.name",
		"r.id AS region_id",
		"r.name AS region_name",
	).
		From("cities AS c").
		Join("regions AS r ON c.region_id = r.id").
		Where(sq.Eq{"c.id": cityIDs}).
		PlaceholderFormat(sq.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query cities: %w", err)
	}
	defer rows.Close()

	cities := make(map[int]models.City)
	for rows.Next() {
		var city models.City
		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.Region.ID,
			&city.Region.Name,
		); err != nil {
			return nil, fmt.Errorf("scan city row: %w", err)
		}

		cities[city.ID] = city
	}

	return cities, nil
}
