package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1OrganizationsGet(ctx context.Context, params api.V1OrganizationsGetParams) (api.V1OrganizationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	organizations, err := h.organizationService.Get(ctx, service.OrganizationGetParams{
		Offset: uint64(params.Offset.Or(0)),
		Limit:  uint64(params.Limit.Or(100)),
	})
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
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

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

	return &api.V1OrganizationsOrganizationIDGetOK{
		Data: models.ConvertOrganizationModelToApi(organization),
	}, nil
}

func (h *Handler) V1OrganizationsContractorsGet(ctx context.Context, params api.V1OrganizationsContractorsGetParams) (api.V1OrganizationsContractorsGetRes, error) {
	organizations, err := h.organizationService.Get(ctx, service.OrganizationGetParams{
		IsContractor: models.Optional[bool]{Value: true, Set: true},
		Offset:       uint64(params.Offset.Or(0)),
		Limit:        uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get organizations: %w", err)
	}

	return &api.V1OrganizationsContractorsGetOK{
		Data: convert.Slice[[]models.Organization, []api.Organization](
			organizations, models.ConvertOrganizationModelToApi),
	}, nil
}
