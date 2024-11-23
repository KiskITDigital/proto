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
	Organization
	Comments
	Suggest
	Verification
	Employee
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
	V1TendersGet(ctx context.Context, params api.V1TendersGetParams) (api.V1TendersGetRes, error)
	V1OrganizationsOrganizationIDTendersGet(ctx context.Context, params api.V1OrganizationsOrganizationIDTendersGetParams) (api.V1OrganizationsOrganizationIDTendersGetRes, error)
	V1TendersTenderIDCommentsGet(ctx context.Context, params api.V1TendersTenderIDCommentsGetParams) (api.V1TendersTenderIDCommentsGetRes, error)
	V1TendersTenderIDCommentsPost(ctx context.Context, req *api.V1TendersTenderIDCommentsPostReq, params api.V1TendersTenderIDCommentsPostParams) (api.V1TendersTenderIDCommentsPostRes, error)
	V1TendersTenderIDRespondPost(ctx context.Context, req *api.V1TendersTenderIDRespondPostReq, params api.V1TendersTenderIDRespondPostParams) (api.V1TendersTenderIDRespondPostRes, error)
	V1TendersVerificationsGet(ctx context.Context, params api.V1TendersVerificationsGetParams) (api.V1TendersVerificationsGetRes, error)
}

type Users interface {
	V1UsersUserIDGet(ctx context.Context, params api.V1UsersUserIDGetParams) (api.V1UsersUserIDGetRes, error)
	V1UsersConfirmEmailPost(ctx context.Context, req *api.V1UsersConfirmEmailPostReq) (api.V1UsersConfirmEmailPostRes, error)
	V1UsersConfirmPasswordPost(ctx context.Context, req *api.V1UsersConfirmPasswordPostReq) (api.V1UsersConfirmPasswordPostRes, error)
	V1UsersGet(ctx context.Context, params api.V1UsersGetParams) (api.V1UsersGetRes, error)
	V1UsersRequestEmailVerificationPost(ctx context.Context, req *api.V1UsersRequestEmailVerificationPostReq) (api.V1UsersRequestEmailVerificationPostRes, error)
	V1UsersRequestResetPasswordPost(ctx context.Context, req *api.V1UsersRequestResetPasswordPostReq) (api.V1UsersRequestResetPasswordPostRes, error)
}

type Survey interface {
	V1SurveyPost(ctx context.Context, req *api.V1SurveyPostReq) (api.V1SurveyPostRes, error)
}

type Catalog interface {
	V1CatalogObjectsGet(ctx context.Context, params api.V1CatalogObjectsGetParams) (api.V1CatalogObjectsGetRes, error)
	V1CatalogServicesGet(ctx context.Context, params api.V1CatalogServicesGetParams) (api.V1CatalogServicesGetRes, error)
	V1CatalogCitiesPost(ctx context.Context, req *api.V1CatalogCitiesPostReq) (api.V1CatalogCitiesPostRes, error)
	V1CatalogRegionsPost(ctx context.Context, req *api.V1CatalogRegionsPostReq) (api.V1CatalogRegionsPostRes, error)
	V1CatalogObjectsPost(ctx context.Context, req *api.V1CatalogObjectsPostReq) (api.V1CatalogObjectsPostRes, error)
	V1CatalogServicesPost(ctx context.Context, req *api.V1CatalogServicesPostReq) (api.V1CatalogServicesPostRes, error)
}

type Organization interface {
	V1OrganizationsGet(ctx context.Context, params api.V1OrganizationsGetParams) (api.V1OrganizationsGetRes, error)
	V1OrganizationsOrganizationIDVerificationsGet(ctx context.Context, params api.V1OrganizationsOrganizationIDVerificationsGetParams) (api.V1OrganizationsOrganizationIDVerificationsGetRes, error)
	V1OrganizationsOrganizationIDVerificationsPost(ctx context.Context, req *api.V1OrganizationsOrganizationIDVerificationsPostReq, params api.V1OrganizationsOrganizationIDVerificationsPostParams) (api.V1OrganizationsOrganizationIDVerificationsPostRes, error)
	V1OrganizationsVerificationsGet(ctx context.Context, params api.V1OrganizationsVerificationsGetParams) (api.V1OrganizationsVerificationsGetRes, error)
}

type Comments interface {
	V1CommentsVerificationsGet(ctx context.Context, params api.V1CommentsVerificationsGetParams) (api.V1CommentsVerificationsGetRes, error)
}

type Suggest interface {
	V1SuggestCompanyGet(ctx context.Context, params api.V1SuggestCompanyGetParams) (api.V1SuggestCompanyGetRes, error)
	V1SuggestCityGet(ctx context.Context, params api.V1SuggestCityGetParams) (api.V1SuggestCityGetRes, error)
}

type Verification interface {
	V1VerificationsRequestIDAprovePost(ctx context.Context, params api.V1VerificationsRequestIDAprovePostParams) (api.V1VerificationsRequestIDAprovePostRes, error)
	V1VerificationsRequestIDDenyPost(ctx context.Context, params api.V1VerificationsRequestIDDenyPostParams) (api.V1VerificationsRequestIDDenyPostRes, error)
}

type Employee interface {
	V1EmployeePost(ctx context.Context, req *api.V1EmployeePostReq) (api.V1EmployeePostRes, error)
}

type RouterParams struct {
	Error        Error
	Auth         Auth
	Tenders      Tenders
	Catalog      Catalog
	Users        Users
	Survey       Survey
	Organization Organization
	Comments     Comments
	Suggest      Suggest
	Verification Verification
	Employee     Employee
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:         params.Auth,
		Error:        params.Error,
		Tenders:      params.Tenders,
		Catalog:      params.Catalog,
		Users:        params.Users,
		Survey:       params.Survey,
		Organization: params.Organization,
		Comments:     params.Comments,
		Suggest:      params.Suggest,
		Verification: params.Verification,
		Employee:     params.Employee,
	}
}
