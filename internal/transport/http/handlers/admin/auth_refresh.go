package admin

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1AdminAuthRefreshPost(ctx context.Context, params api.V1AdminAuthRefreshPostParams) (api.V1AdminAuthRefreshPostRes, error) {
	res, err := h.svc.Refresh(ctx, params.UbratoAdminSession)
	if err != nil {
		return nil, fmt.Errorf("refresh session: %w", err)
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

	return &api.V1AdminAuthRefreshPostOKHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AdminAuthRefreshPostOK{
			Data: api.V1AdminAuthRefreshPostOKData{
				User:        models.ConvertAdminUserModelToApi(res.User),
				AccessToken: res.AccessToken,
			},
		},
	}, nil
}
