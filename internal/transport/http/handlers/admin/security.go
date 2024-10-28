package admin

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) HandleAdminCookieAuth(ctx context.Context, operationName string, t api.AdminCookieAuth) (context.Context, error) {
	return ctx, nil
}

func (h *Handler) HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error) {
	claims, err := h.svc.ValidateAccessToken(ctx, t.GetToken())
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, models.UserIDKey, claims.UserID)
	ctx = context.WithValue(ctx, models.RoleKey, claims.Role)

	return ctx, nil
}
