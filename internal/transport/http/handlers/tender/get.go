package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDGet(ctx context.Context, params api.V1TendersTenderIDGetParams) (api.V1TendersTenderIDGetRes, error) {
	tender, err := h.svc.GetByID(ctx, params.TenderID)
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersTenderIDGetOK{
		Data: api.V1TendersTenderIDGetOKData{
			Tender: models.ConvertTenderModelToApi(tender),
		},
	}, nil
}

func (h *Handler) V1TendersGet(ctx context.Context) (api.V1TendersGetRes, error) {
	tenders, err := h.svc.Get(ctx, service.TenderGetParams{})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersGetOK{
		Data: api.V1TendersGetOKData{
			Tenders: convert.Slice[[]models.Tender, []api.Tender](tenders, models.ConvertTenderModelToApi),
		},
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDTendersGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDTendersGetParams,
) (api.V1OrganizationsOrganizationIDTendersGetRes, error) {
	organizationID := contextor.GetOrganizationID(ctx)

	tenders, err := h.svc.Get(ctx, service.TenderGetParams{
		OrganizationID: models.Optional[int]{Value: params.OrganizationID, Set: true},
		WithDrafts:     organizationID == params.OrganizationID,
	})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDTendersGetOK{
		Data: api.V1OrganizationsOrganizationIDTendersGetOKData{
			Tenders: convert.Slice[[]models.Tender, []api.Tender](tenders, models.ConvertTenderModelToApi),
		},
	}, nil
}
