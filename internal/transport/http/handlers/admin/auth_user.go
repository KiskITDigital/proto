package admin

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service/admin"
)

func (h *Handler) V1AdminAuthUserGet(ctx context.Context) (api.V1AdminAuthUserGetRes, error) {
	res, err := h.svc.GetUser(ctx, admin.GetUserParams{
		ID: contextor.GetUserID(ctx),
	})
	if err != nil {
		return nil, err
	}

	return &api.V1AdminAuthUserGetOK{
		Data: api.V1AdminAuthUserGetOKData{
			User: models.ConvertAdminUserModelToApi(res.User),
		},
	}, nil
}
