package tender

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) GetObjects(ctx context.Context, qe store.QueryExecutor, objectIDs []int) (map[int]models.Object, error) {
	query := `WITH RECURSIVE object_hierarchy AS (
		SELECT 
			id,
			name,
			parent_id
		FROM 
			objects
		WHERE 
			id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			o.id,
			o.name,
			o.parent_id
		FROM 
			objects o
		JOIN 
			object_hierarchy sh ON o.id = sh.parent_id
	)

	SELECT
		id,
		parent_id,
		name
	FROM 
		object_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(objectIDs))
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	objects := map[int]models.Object{}

	for rows.Next() {
		var (
			object   models.Object
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&object.ID,
			&parentID,
			&object.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		object.ParentID = int(parentID.Int64)

		objects[object.ID] = object
	}

	return objects, nil
}
