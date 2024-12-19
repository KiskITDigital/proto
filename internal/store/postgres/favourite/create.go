package favourite

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *FavouriteStore) Create(ctx context.Context, qe store.QueryExecutor, params store.FavouriteCreateParams) error {
	builder := squirrel.Insert("favourites").Columns(
		"organization_id",
		"object_type",
		"object_id",
	).Values(
		params.OrganizationID,
		params.ObjectType,
		params.ObjectID,
	).PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
