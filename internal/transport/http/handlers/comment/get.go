package comment

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

func (h *Handler) V1CommentsVerificationsGet(ctx context.Context, params api.V1CommentsVerificationsGetParams) (api.V1CommentsVerificationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeComment,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1CommentsVerificationsGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}
