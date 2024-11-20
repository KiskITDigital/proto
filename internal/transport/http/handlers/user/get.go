package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
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

func (h *Handler) V1UsersGet(ctx context.Context, params api.V1UsersGetParams) (api.V1UsersGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	users, err := h.svc.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return &api.V1UsersGetOK{
		Data: convert.Slice[[]models.User, []api.User](users, models.ConvertUserModelToApi),
	}, nil
}
