package admin

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service/admin"
)

func (h *Handler) V1AdminUsersGet(ctx context.Context) (api.V1AdminUsersGetRes, error) {
	res, err := h.svc.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &api.V1AdminUsersGetOK{
		Data: api.V1AdminUsersGetOKData{
			Users: convert.Slice[[]models.AdminUser, []api.AdminUser](res.Users, models.ConvertAdminUserModelToApi),
		},
	}, nil
}

func (h *Handler) V1AdminUsersPost(ctx context.Context, req *api.V1AdminUsersPostReq) (api.V1AdminUsersPostRes, error) {
	res, err := h.svc.CreateUser(ctx, admin.CreateUserParams{
		Email:      string(req.Email),
		Phone:      string(req.Phone),
		Password:   string(req.Password),
		FirstName:  string(req.FirstName),
		LastName:   string(req.LastName),
		MiddleName: string(req.MiddleName),
		AvatarURL:  string(req.AvatarURL.Value),
		Role:       models.UserRole(req.Role),
	})
	if err != nil {
		return nil, err
	}

	return &api.V1AdminUsersPostCreated{
		Data: api.V1AdminUsersPostCreatedData{
			User: models.ConvertAdminUserModelToApi(res.User),
		},
	}, nil
}

func (h *Handler) V1AdminUsersUserIDGet(ctx context.Context, params api.V1AdminUsersUserIDGetParams) (api.V1AdminUsersUserIDGetRes, error) {
	res, err := h.svc.GetUser(ctx, admin.GetUserParams{
		ID: params.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &api.V1AdminUsersUserIDGetOK{
		Data: api.V1AdminUsersUserIDGetOKData{
			User: models.ConvertAdminUserModelToApi(res.User),
		},
	}, nil
}
