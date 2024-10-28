package admin

import (
	"context"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service/admin"
)

func (h *Handler) V1AdminAuthSigninPost(ctx context.Context, req *api.V1AdminAuthSigninPostReq) (api.V1AdminAuthSigninPostRes, error) {
	res, err := h.svc.SignIn(ctx, admin.SignInParams{
		Email:    string(req.Email),
		Password: string(req.Password),
	})
	if err != nil {
		return nil, err
	}

	cookie := http.Cookie{
		Name:     "ubrato_admin_session",
		Value:    res.Session.ID,
		Path:     "/",
		Expires:  res.Session.ExpiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return &api.V1AdminAuthSigninPostOKHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AdminAuthSigninPostOK{
			Data: api.V1AdminAuthSigninPostOKData{
				User:        models.ConvertAdminUserModelToApi(res.User),
				AccessToken: res.AccessToken,
			},
		},
	}, nil
}
