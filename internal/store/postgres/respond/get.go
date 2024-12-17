package respond

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *RespondStore) Get(ctx context.Context, qe store.QueryExecutor, params store.RespondGetParams) ([]models.Respond, error) {
	builder := squirrel.Select(
		"r.tender_id",
		"r.organization_id",
		"r.price",
		"r.is_nds_price",
		"r.created_at",
	).From("tender_responses AS r").
		Where(squirrel.Eq{"r.tender_id": params.TenderID}).
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	responds := []models.Respond{}
	for rows.Next() {
		var respond models.Respond

		if err = rows.Scan(
			&respond.TenderID,
			&respond.OrganizationID,
			&respond.Price,
			&respond.IsNDSPrice,
			&respond.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		responds = append(responds, respond)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return responds, nil
}
