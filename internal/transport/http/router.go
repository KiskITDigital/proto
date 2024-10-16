package http

import (
	"context"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

var _ api.Handler = new(Router)

type Router struct {
	Auth
	Error
}

type Error interface {
	HandleError(ctx context.Context, w http.ResponseWriter, _ *http.Request, err error)
}

type Auth interface {
	V1AuthSigninPost(ctx context.Context, req *api.V1AuthSigninPostReq) (api.V1AuthSigninPostRes, error)
	V1AuthSignupPost(ctx context.Context, req *api.V1AuthSignupPostReq) (api.V1AuthSignupPostRes, error)
}

type RouterParams struct {
	Error Error
	Auth  Auth
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:  params.Auth,
		Error: params.Error,
	}
}
