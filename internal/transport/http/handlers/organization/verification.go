package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1OrganizationsOrganizationIDVerificationsPost(
	ctx context.Context,
	req []api.Attachment,
	params api.V1OrganizationsOrganizationIDVerificationsPostParams,
) (api.V1OrganizationsOrganizationIDVerificationsPostRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to request verification of the organization", nil)
	}

	if err := h.organizationService.CreateVerificationRequest(ctx, service.OrganizationCreateVerificationRequestParams{
		OrganizationID: params.OrganizationID,
		Attachments:    convert.Slice[[]api.Attachment, []models.Attachment](req, models.ConvertAPIToAttachment),
	}); err != nil {
		return nil, fmt.Errorf("create verif req: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDVerificationsPostOK{}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDVerificationsGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDVerificationsGetParams,
) (api.V1OrganizationsOrganizationIDVerificationsGetRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to get verification requests of the organization", nil)
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeOrganization,
		ObjectID:   models.NewOptional(params.OrganizationID)})
	if err != nil {
		return nil, fmt.Errorf("get verif requests: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDVerificationsGetOK{
		Data: convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		// В будущем может нужно будет
		// Pagination: pagination.ConvertPaginationToAPI(request.Pagination),
	}, nil
}

func (h *Handler) V1OrganizationsVerificationsGet(ctx context.Context, params api.V1OrganizationsVerificationsGetParams) (api.V1OrganizationsVerificationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeOrganization,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1OrganizationsVerificationsGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}
