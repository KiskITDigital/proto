package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1UsersRequestEmailVerificationPost(ctx context.Context, req *api.V1UsersRequestEmailVerificationPostReq) (api.V1UsersRequestEmailVerificationPostRes, error) {
	err := h.svc.ReqEmailVerification(ctx, string(req.GetEmail()))
	if err != nil {
		return nil, err
	}
	
	return &api.V1UsersRequestEmailVerificationPostOK{}, nil
}

func (h *Handler) V1UsersConfirmEmailPost(ctx context.Context, req *api.V1UsersConfirmEmailPostReq) (api.V1UsersConfirmEmailPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
