package tender

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

func (h *Handler) V1TendersTenderIDGet(ctx context.Context, params api.V1TendersTenderIDGetParams) (api.V1TendersTenderIDGetRes, error) {
	tender, err := h.tenderService.GetByID(ctx, params.TenderID)
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersTenderIDGetOK{
		Data: models.ConvertTenderModelToApi(tender),
	}, nil
}

func (h *Handler) V1TendersGet(ctx context.Context, params api.V1TendersGetParams) (api.V1TendersGetRes, error) {
	tenders, err := h.tenderService.List(ctx, service.TenderListParams{})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersGetOK{
		Data: convert.Slice[[]models.Tender, []api.Tender](tenders, models.ConvertTenderModelToApi),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDTendersGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDTendersGetParams,
) (api.V1OrganizationsOrganizationIDTendersGetRes, error) {
	organizationID := contextor.GetOrganizationID(ctx)

	tenders, err := h.tenderService.List(ctx, service.TenderListParams{
		OrganizationID: models.Optional[int]{Value: params.OrganizationID, Set: true},
		WithDrafts:     organizationID == params.OrganizationID,
	})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDTendersGetOK{
		Data: convert.Slice[[]models.Tender, []api.Tender](tenders, models.ConvertTenderModelToApi),
	}, nil
}

func (h *Handler) V1TendersVerificationsGet(ctx context.Context, params api.V1TendersVerificationsGetParams) (api.V1TendersVerificationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	orgVerifications, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeTender,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Offset:     uint64(params.Offset.Or(0)),
		Limit:      uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1TendersVerificationsGetOK{
		Data: convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](
			orgVerifications, models.VerificationRequestModelToApi),
	}, nil
}
