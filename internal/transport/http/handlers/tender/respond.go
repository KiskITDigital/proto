package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDRespondPost(ctx context.Context, req *api.V1TendersTenderIDRespondPostReq, params api.V1TendersTenderIDRespondPostParams) (api.V1TendersTenderIDRespondPostRes, error) {
	err := h.tenderService.Respond(ctx, service.TenderRespondParams{
		TenderID:       params.TenderID,
		OrganizationID: contextor.GetOrganizationID(ctx),
		Price:          req.Price,
		IsNds:          req.IsNds,
	})
	if err != nil {
		return nil, fmt.Errorf("respond: %w", err)
	}

	return &api.V1TendersTenderIDRespondPostOK{}, nil
}
