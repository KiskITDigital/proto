package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type OrganizationGetParams struct {}

func (s *Service) Get(ctx context.Context, params OrganizationGetParams) ([]models.Organization, error) {
	return s.organizationStore.Get(ctx, s.psql.DB(), store.OrganizationGetParams{})
}
