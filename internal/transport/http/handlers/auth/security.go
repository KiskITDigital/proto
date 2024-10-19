package auth

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
	return ctx, nil
}
