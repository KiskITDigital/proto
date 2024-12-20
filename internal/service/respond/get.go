package respond

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.RespondGetParams) (models.RespondPagination, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		return models.RespondPagination{}, fmt.Errorf("get tender by id: %w", err)
	}

	if tender.Organization.ID != contextor.GetOrganizationID(ctx) {
		return models.RespondPagination{}, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to get the responds", nil)
	}

	responds, err := s.respondStore.Get(ctx, s.psql.DB(), store.RespondGetParams{
		TenderID: params.TenderID,
		Offset:   models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0},
		Limit:    models.NewOptional(params.PerPage)})
	if err != nil {
		return models.RespondPagination{}, fmt.Errorf("get respond: %w", err)
	}

	count, err := s.respondStore.Count(ctx, s.psql.DB(), store.RespondGetCountParams{
		TenderID: params.TenderID})
	if err != nil {
		return models.RespondPagination{}, fmt.Errorf("get count responds: %w", err)
	}

	return models.RespondPagination{
		Responds:   responds,
		Pagination: pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}
