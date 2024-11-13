package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) List(ctx context.Context, params service.TenderListParams) ([]models.Tender, error) {
	return s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
		OrganizationID: params.OrganizationID,
		WithDrafts:     params.WithDrafts,
	})
}
