package verification

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestUpdateStatusParams) (store.VerificationObjectUpdateStatusResult, error) {
	builder := squirrel.Update("verification_requests").
		Set("status", params.Status).
		Set("reviewed_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Suffix(`
		returning 
			object_id,
			object_type`).
		Where(squirrel.Eq{"id": params.ID}).
		PlaceholderFormat(squirrel.Dollar)

	result := store.VerificationObjectUpdateStatusResult{}
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&result.ObjectID,
		&result.ObjectType,
	); err != nil {
		return result, fmt.Errorf("query row: %w", err)
	}

	return result, nil
}
