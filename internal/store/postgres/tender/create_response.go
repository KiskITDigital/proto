package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) CreateResponse(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateResponseParams) error {
	builder := squirrel.
		Insert("tenders_responses").
		Columns(
			"tender_id",
			"organization_id",
			"price",
			"is_nds",
		).
		Values(
			params.TenderID,
			params.OrganizationID,
			params.Price,
			params.IsNds,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
