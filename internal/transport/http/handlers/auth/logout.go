package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1AuthLogoutDelete(ctx context.Context, params api.V1AuthLogoutDeleteParams) (api.V1AuthLogoutDeleteRes, error) {
	if err := h.authSvc.Logout(ctx, params.UbratoSession); err != nil {
		return nil, fmt.Errorf("logout: %w", err)
	}

	cookie := &http.Cookie{
		Name:     "ubrato_session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	return &api.V1AuthLogoutDeleteNoContent{
		SetCookie: api.NewOptString(cookie.String()),
	}, nil
}
