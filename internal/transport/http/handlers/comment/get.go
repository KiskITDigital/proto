package comment

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

func (h *Handler) V1CommentsVerificationsGet(ctx context.Context, params api.V1CommentsVerificationsGetParams) (api.V1CommentsVerificationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	commentVerifications, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeComment,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Offset:     uint64(params.Offset.Or(0)),
		Limit:      uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1CommentsVerificationsGetOK{
		Data: convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](
			commentVerifications, models.VerificationRequestModelToApi),
	}, nil
}
