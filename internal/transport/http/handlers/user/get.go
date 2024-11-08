package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1UsersUserIDGet(ctx context.Context, params api.V1UsersUserIDGetParams) (api.V1UsersUserIDGetRes, error) {
	user, err := h.svc.GetByID(ctx, params.UserID)
	if err != nil {
		return nil, fmt.Errorf("get user with id: %d", params.UserID)
	}

	return &api.V1UsersUserIDGetOK{
		Data: models.ConvertUserModelToApi(user),
	}, nil
}
