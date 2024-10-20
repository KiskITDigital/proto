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
	Tenders
}

type Error interface {
	HandleError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)
}

type Auth interface {
	V1AuthSigninPost(ctx context.Context, req *api.V1AuthSigninPostReq) (api.V1AuthSigninPostRes, error)
	V1AuthSignupPost(ctx context.Context, req *api.V1AuthSignupPostReq) (api.V1AuthSignupPostRes, error)
	V1AuthUserGet(ctx context.Context) (api.V1AuthUserGetRes, error)
	V1AuthRefreshPost(ctx context.Context, params api.V1AuthRefreshPostParams) (api.V1AuthRefreshPostRes, error)

	HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error)
	HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error)
}

type Tenders interface {
	V1TendersCreatePost(ctx context.Context, req *api.V1TendersCreatePostReq) (api.V1TendersCreatePostRes, error)
}

type RouterParams struct {
	Error   Error
	Auth    Auth
	Tenders Tenders
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:    params.Auth,
		Error:   params.Error,
		Tenders: params.Tenders,
	}
}
