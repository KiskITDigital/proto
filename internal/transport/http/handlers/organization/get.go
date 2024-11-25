package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	organizationService "gitlab.ubrato.ru/ubrato/core/internal/service/organization"
)

func (h *Handler) V1OrganizationsGet(ctx context.Context, params api.V1OrganizationsGetParams) (api.V1OrganizationsGetRes, error) {
	organizations, err := h.organizationService.Get(ctx, organizationService.OrganizationGetParams{})
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

func (h *Handler) V1OrganizationsVerificationsGet(ctx context.Context, params api.V1OrganizationsVerificationsGetParams) (api.V1OrganizationsVerificationsGetRes, error) {
	orgVerifications, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeOrganization,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Offset:     uint64(params.Offset.Or(0)),
		Limit:      uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1OrganizationsVerificationsGetOK{
		Data: convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](
			orgVerifications, models.VerificationRequestModelToApi),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDGet(ctx context.Context, params api.V1OrganizationsOrganizationIDGetParams) (api.V1OrganizationsOrganizationIDGetRes, error) {
	// TODO: check role
	organization, err := h.organizationService.GetByID(ctx, params.OrganizationID)
	if err != nil {
		return nil, fmt.Errorf("get organization by id: %w", err)
	}

	apiOrganization := models.ConvertOrganizationModelToApi(organization)
	return &apiOrganization, nil
}
