package tender

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) GetServices(ctx context.Context, qe store.QueryExecutor, serviceIDs []int) (map[int]models.Service, error) {
	query := `WITH RECURSIVE service_hierarchy AS (
		SELECT 
			id,
			parent_id,
			name
		FROM 
			services
		WHERE 
			id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			s.id,
			s.parent_id,
			s.name
		FROM 
			services s
		JOIN 
			service_hierarchy sh ON s.id = sh.parent_id
	)

	SELECT
		id,
		parent_id,
		name
	FROM
		service_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(serviceIDs))
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	services := map[int]models.Service{}

	for rows.Next() {
		var (
			service  models.Service
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&service.ID,
			&parentID,
			&service.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		service.ParentID = int(parentID.Int64)

		services[service.ID] = service
	}

	return services, nil
}
