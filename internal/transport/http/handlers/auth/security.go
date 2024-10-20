package auth

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
	return ctx, nil
}

func (h *Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	claims, err := h.svc.ValidateAccessToken(ctx, t.GetToken())
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, models.AccessTokenKey, claims)

	return ctx, nil
}
