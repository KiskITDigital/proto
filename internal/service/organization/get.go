package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.OrganizationGetParams) ([]models.Organization, error) {
	return s.organizationStore.Get(ctx, s.psql.DB(), store.OrganizationGetParams{
		IsContractor: params.IsContractor,
		Limit:        params.Limit,
		Offset:       params.Offset,
	})
}

func (s *Service) GetByID(ctx context.Context, id int) (models.Organization, error) {
	return s.organizationStore.GetByID(ctx, s.psql.DB(), id)
}
