package auth

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1AuthSigninPost(ctx context.Context, req *api.V1AuthSigninPostReq) (api.V1AuthSigninPostRes, error) {
	return &api.V1AuthSigninPostOKHeaders{}, nil
}
