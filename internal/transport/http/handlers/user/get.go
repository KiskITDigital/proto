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
		return nil, fmt.Errorf("get user by id")
	}

	return &api.V1UsersUserIDGetOK{
		Data: api.V1UsersUserIDGetOKData{
			User: models.ConvertUserModelToApi(user),
		},
	}, nil
}
