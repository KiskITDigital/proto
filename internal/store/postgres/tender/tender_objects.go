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

func (s *TenderStore) GetTendersObjects(ctx context.Context, qe store.QueryExecutor, tenderIDs []int) ([]models.TenderObject, error) {
	query := `WITH RECURSIVE object_hierarchy AS (
		SELECT 
			o.id AS object_id,
			o.name AS object_name,
			o.parent_id,
			tobj.tender_id
		FROM 
			objects o
		JOIN 
			tender_objects tobj ON o.id = tobj.object_id
		WHERE 
			tobj.tender_id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			o.id AS object_id,
			o.name AS object_name,
			o.parent_id,
			sh.tender_id
		FROM 
			objects o
		JOIN 
			object_hierarchy sh ON o.id = sh.parent_id
	)

	SELECT DISTINCT ON (object_id, tender_id)
		object_id,
		object_name,
		parent_id,
		tender_id
	FROM 
		object_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(tenderIDs))
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
			&object.TenderID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		object.ParentID = int(parentID.Int64)

		objects = append(objects, object)
	}

	return objects, nil
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

func (s *TenderStore) DeleteTenderObjects(ctx context.Context, qe store.QueryExecutor, params store.TenderObjectsDeleteParams) error {
	builder := squirrel.
		Delete("tender_objects").
		Where(squirrel.Eq{"tender_id": params.TenderID}).
		Where(squirrel.Eq{"service_id": params.ObjectsIDs}).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
