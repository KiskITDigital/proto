package auth

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1AuthUserGet(ctx context.Context) (api.V1AuthUserGetRes, error) {
	return &api.V1AuthUserGetCreatedHeaders{}, nil
}
