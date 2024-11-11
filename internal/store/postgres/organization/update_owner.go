package organization

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) UpdateOwner(ctx context.Context, qe store.QueryExecutor, organizationID, ownerID int) error {
	builder := squirrel.
		Update("organizations").
		Set("owner_user_id", ownerID).
		Where(squirrel.Eq{"id": organizationID})

	_, err := builder.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
