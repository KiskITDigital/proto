package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

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
