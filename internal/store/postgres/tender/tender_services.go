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

func (s *TenderStore) GetTendersServices(ctx context.Context, qe store.QueryExecutor, tenderIDs []int) ([]models.TenderService, error) {
	query := `WITH RECURSIVE service_hierarchy AS (
		SELECT 
			s.id AS service_id,
			s.name AS service_name,
			s.parent_id,
			ts.tender_id
		FROM 
			services s
		JOIN 
			tender_services ts ON s.id = ts.service_id
		WHERE 
			ts.tender_id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			s.id AS service_id,
			s.name AS service_name,
			s.parent_id,
			sh.tender_id
		FROM 
			services s
		JOIN 
			service_hierarchy sh ON s.id = sh.parent_id
	)

	SELECT DISTINCT ON (service_id, tender_id)
		service_id,
		service_name,
		parent_id,
		tender_id
	FROM 
		service_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(tenderIDs))
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
			&service.TenderID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		service.ParentID = int(parentID.Int64)

		services = append(services, service)
	}

	return services, nil
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

func (s *TenderStore) DeleteTenderServices(ctx context.Context, qe store.QueryExecutor, params store.TenderServicesDeleteParams) error {
	builder := squirrel.
		Delete("tender_services").
		Where(squirrel.Eq{"tender_id": params.TenderID}).
		Where(squirrel.Eq{"service_id": params.ServicesIDs}).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
