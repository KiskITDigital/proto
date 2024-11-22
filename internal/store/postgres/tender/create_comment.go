package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) CreateComment(ctx context.Context, qe store.QueryExecutor, params store.CommentCreateParams) error {
	builder := squirrel.
		Insert("comments").
		Columns(
			"organization_id",
			"object_type",
			"object_id",
			"content",
			"verification_status",
			"attachments",
		).
		Values(
			params.OrganizationID,
			models.ObjectTypeTender,
			params.ObjectID,
			params.Content,
			params.Attachments,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
