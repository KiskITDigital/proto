package verification

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1VerificationsRequestIDGet(ctx context.Context, params api.V1VerificationsRequestIDGetParams) (api.V1VerificationsRequestIDGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	request, err := h.svc.GetByID(ctx, params.RequestID)
	if err != nil {
		return nil, fmt.Errorf("get by id: %w", err)
	}

	return &api.V1VerificationsRequestIDGetOK{
		Data: models.VerificationRequestModelToApi(request),
	}, nil
}
