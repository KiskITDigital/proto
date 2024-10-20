package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

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
