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
		Data: models.ConvertRegularUserModelToApi(user),
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
		Data: convert.Slice[[]models.FullUser, []api.V1UsersGetOKDataItem](users, func(fu models.FullUser) api.V1UsersGetOKDataItem {
			var user api.V1UsersGetOKDataItem

			if fu.Role != 0 {
				fu.EmployeeUser.User = fu.User
				user.Type = api.EmployeeUserV1UsersGetOKDataItem
				user.EmployeeUser = models.ConvertEmployeeUserModelToApi(fu.EmployeeUser)
				return user
			}

			fu.RegularUser.User = fu.User
			user.Type = api.RegularUserV1UsersGetOKDataItem
			user.RegularUser = models.ConvertRegularUserModelToApi(fu.RegularUser)
			return user
		}),
	}, nil
}
