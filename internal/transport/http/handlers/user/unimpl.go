package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1UsersConfirmEmailPost(ctx context.Context, req *api.V1UsersConfirmEmailPostReq) (api.V1UsersConfirmEmailPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1UsersConfirmPasswordPost(ctx context.Context, req *api.V1UsersConfirmPasswordPostReq) (api.V1UsersConfirmPasswordPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1UsersGet(ctx context.Context, params api.V1UsersGetParams) (api.V1UsersGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1UsersRequestEmailVerificationPost(ctx context.Context, req *api.V1UsersRequestEmailVerificationPostReq) (api.V1UsersRequestEmailVerificationPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1UsersRequestResetPasswordPost(ctx context.Context, req *api.V1UsersRequestResetPasswordPostReq) (api.V1UsersRequestResetPasswordPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
