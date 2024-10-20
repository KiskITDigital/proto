package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
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
