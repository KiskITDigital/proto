package verification

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error {
	builder := squirrel.
		Insert("verification_requests").
		Columns(
			"object_type",
			"object_id",
			"attachments",
		).
		Values(
			params.ObjectType,
			params.ObjectType,
			pq.Array(params.Attachments),
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
