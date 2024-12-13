package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetCustomer(ctx context.Context, organizationId int) (models.Organization, error) {
	return s.organizationStore.GetCustomer(ctx, s.psql.DB(), organizationId)
}

func (s *Service) GetContractor(ctx context.Context, organizationId int) (models.Organization, error) {
	return s.organizationStore.GetContractor(ctx, s.psql.DB(), organizationId)
}
