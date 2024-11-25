package comment

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CommentStore) CreateComment(ctx context.Context, qe store.QueryExecutor, params store.CommentCreateParams) error {
	builder := squirrel.
		Insert("comments").
		Columns(
			"organization_id",
			"object_type",
			"object_id",
			"title",
			"content",
			"verification_status",
			"attachments",
		).
		Values(
			params.OrganizationID,
			params.ObjectType,
			params.ObjectID,
			params.Title,
			params.Content,
			params.VerificationStatus,
			pq.Array(params.Attachments),
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
