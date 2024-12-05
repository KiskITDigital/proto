package organization

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1OrganizationsOrganizationIDProfileBrandPut(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDProfileBrandPutReq,
	params api.V1OrganizationsOrganizationIDProfileBrandPutParams,
) (api.V1OrganizationsOrganizationIDProfileBrandPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	if err := h.organizationService.UpdateBrand(ctx, service.OrganizationUpdateBrandParams{
		OrganizationID: params.OrganizationID,
		Brand:          models.Optional[string]{Value: req.GetBrand().Value, Set: req.GetBrand().Set},
		AvatarURL:      models.Optional[string]{Value: string(req.GetAvatarURL().Value), Set: req.GetAvatarURL().Set},
	}); err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileBrandPutOK{}, nil
}
