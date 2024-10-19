package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1AuthRefreshPost(ctx context.Context, params api.V1AuthRefreshPostParams) (api.V1AuthRefreshPostRes, error) {
	resp, err := h.svc.Refresh(ctx, params.UbratoSession)
	if err != nil {
		return nil, fmt.Errorf("refresh session: %w", err)
	}

	cookie := http.Cookie{
		Name:     "ubrato_session",
		Value:    resp.Session.ID,
		Path:     "/",
		Expires:  resp.Session.ExpiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return &api.V1AuthRefreshPostOKHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AuthRefreshPostOK{
			Data: api.V1AuthRefreshPostOKData{
				User:        ConvertUserModelToApi(resp.User),
				AccessToken: resp.AccessToken,
			},
		},
	}, nil
}
