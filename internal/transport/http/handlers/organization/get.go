package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	organizationService "gitlab.ubrato.ru/ubrato/core/internal/service/organization"
)

func (h *Handler) V1OrganizationsGet(ctx context.Context) (api.V1OrganizationsGetRes, error) {
	organizations, err := h.svc.Get(ctx, organizationService.OrganizationGetParams{})
	if err != nil {
		return nil, fmt.Errorf("get organizations: %w", err)
	}

	return &api.V1OrganizationsGetOK{
		Data: api.V1OrganizationsGetOKData{
			Organizations: convert.Slice[[]models.Organization, []api.Organization](
				organizations, models.ConvertOrganizationModelToApi),
		},
	}, nil
}
