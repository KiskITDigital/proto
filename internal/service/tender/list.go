package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) List(ctx context.Context, params service.TenderListParams) (models.TendersRes, error) {
	tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
		OrganizationID: params.OrganizationID,
		VerifiedOnly:   params.VerifiedOnly,
		WithDrafts:     params.WithDrafts,
		Limit:          models.Optional[uint64]{Value: params.PerPage, Set: params.PerPage != 0},
		Offset:         models.Optional[uint64]{Value: params.Page * params.PerPage, Set: params.Page != 0}})
	if err != nil {
		return models.TendersRes{}, fmt.Errorf("get tenders: %w", err)
	}

	count, err := s.tenderStore.Count(ctx, s.psql.DB(), store.TenderGetCountParams{
		OrganizationID: params.OrganizationID,
		VerifiedOnly:   params.VerifiedOnly,
		WithDrafts:     params.WithDrafts})
	if err != nil {
		return models.TendersRes{}, fmt.Errorf("get count tenders: %w", err)
	}

	return models.TendersRes{
		Tenders: tenders,
		Pagination: models.Pagination{
			Page:    params.Page,
			Pages:   pagination.CalculatePages(count, params.PerPage),
			PerPage: params.PerPage,
		},
	}, nil
}
