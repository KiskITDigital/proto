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
	Catalog
	Users
	Survey
	Admin
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
	V1TendersPost(ctx context.Context, req *api.V1TendersPostReq) (api.V1TendersPostRes, error)
	V1TendersTenderIDPut(ctx context.Context, req *api.V1TendersTenderIDPutReq, params api.V1TendersTenderIDPutParams) (api.V1TendersTenderIDPutRes, error)
	V1TendersTenderIDGet(ctx context.Context, params api.V1TendersTenderIDGetParams) (api.V1TendersTenderIDGetRes, error)
	V1TendersGet(ctx context.Context) (api.V1TendersGetRes, error)
	V1OrganizationsOrganizationIDTendersGet(ctx context.Context, params api.V1OrganizationsOrganizationIDTendersGetParams) (api.V1OrganizationsOrganizationIDTendersGetRes, error)
}

type Users interface {
	V1UsersUserIDGet(ctx context.Context, params api.V1UsersUserIDGetParams) (api.V1UsersUserIDGetRes, error)
}

type Survey interface {
	V1SurveyPost(ctx context.Context, req *api.V1SurveyPostReq) (api.V1SurveyPostRes, error)
}

type Catalog interface {
	V1CatalogObjectsGet(ctx context.Context) (api.V1CatalogObjectsGetRes, error)
	V1CatalogServicesGet(ctx context.Context) (api.V1CatalogServicesGetRes, error)
	V1CatalogCitiesPost(ctx context.Context, req *api.V1CatalogCitiesPostReq) (api.V1CatalogCitiesPostRes, error)
	V1CatalogRegionsPost(ctx context.Context, req *api.V1CatalogRegionsPostReq) (api.V1CatalogRegionsPostRes, error)
	V1CatalogObjectsPost(ctx context.Context, req *api.V1CatalogObjectsPostReq) (api.V1CatalogObjectsPostRes, error)
	V1CatalogServicesPost(ctx context.Context, req *api.V1CatalogServicesPostReq) (api.V1CatalogServicesPostRes, error)
}

type Admin interface {
	V1AdminAuthRefreshPost(ctx context.Context, params api.V1AdminAuthRefreshPostParams) (api.V1AdminAuthRefreshPostRes, error)
	V1AdminAuthSigninPost(ctx context.Context, req *api.V1AdminAuthSigninPostReq) (api.V1AdminAuthSigninPostRes, error)
	V1AdminAuthUserGet(ctx context.Context) (api.V1AdminAuthUserGetRes, error)
	V1AdminUsersGet(ctx context.Context) (api.V1AdminUsersGetRes, error)
	V1AdminUsersPost(ctx context.Context, req *api.V1AdminUsersPostReq) (api.V1AdminUsersPostRes, error)
	V1AdminUsersUserIDGet(ctx context.Context, params api.V1AdminUsersUserIDGetParams) (api.V1AdminUsersUserIDGetRes, error)

	HandleAdminCookieAuth(ctx context.Context, operationName string, t api.AdminCookieAuth) (context.Context, error)
	HandleAdminBearerAuth(ctx context.Context, operationName string, t api.AdminBearerAuth) (context.Context, error)
}

type RouterParams struct {
	Error   Error
	Auth    Auth
	Tenders Tenders
	Catalog Catalog
	Users   Users
	Survey  Survey
	Admin   Admin
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:    params.Auth,
		Error:   params.Error,
		Tenders: params.Tenders,
		Catalog: params.Catalog,
		Users:   params.Users,
		Survey:  params.Survey,
		Admin:   params.Admin,
	}
}
