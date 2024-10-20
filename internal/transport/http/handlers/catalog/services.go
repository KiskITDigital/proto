package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1CatalogServicesGet(ctx context.Context) (api.V1CatalogServicesGetRes, error) {
	services, err := h.svc.GetServices(ctx)
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return &api.V1CatalogServicesGetOK{
		Data: api.V1CatalogServicesGetOKData{
			Services: convert.Slice[models.CatalogServices, api.Services](services, models.ConvertModelCatalogServiceToApi),
		},
	}, nil
}
