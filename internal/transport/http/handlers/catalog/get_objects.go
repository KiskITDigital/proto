package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1CatalogObjectsGet(ctx context.Context) (api.V1CatalogObjectsGetRes, error) {
	objects, err := h.svc.GetObjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return &api.V1CatalogObjectsGetOK{
		Data: api.V1CatalogObjectsGetOKData{
			Objects: convert.Slice[models.CatalogObjects, api.Objects](objects, models.ConvertModelCatalogObjectToApi),
		},
	}, nil
}
