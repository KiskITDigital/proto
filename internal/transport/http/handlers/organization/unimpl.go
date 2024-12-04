package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

// Верификация
func (h *Handler) V1OrganizationsOrganizationIDVerificationsGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDVerificationsGetParams,
) (api.V1OrganizationsOrganizationIDVerificationsGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDVerificationsPost(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDVerificationsPostReq,
	params api.V1OrganizationsOrganizationIDVerificationsPostParams,
) (api.V1OrganizationsOrganizationIDVerificationsPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

// Профиль
func (h *Handler) V1OrganizationsProfileBrandPut(ctx context.Context, req *api.V1OrganizationsProfileBrandPutReq) (api.V1OrganizationsProfileBrandPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsProfileContactsPut(ctx context.Context, req *api.V1OrganizationsProfileContactsPutReq) (api.V1OrganizationsProfileContactsPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsProfileContractorPut(ctx context.Context, req *api.V1OrganizationsProfileContractorPutReq) (api.V1OrganizationsProfileContractorPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsProfileCustomerPut(ctx context.Context, req *api.V1OrganizationsProfileCustomerPutReq) (api.V1OrganizationsProfileCustomerPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

// Портфолио
func (h *Handler) V1OrganizationsPortfolioPortfolioIDDelete(ctx context.Context, params api.V1OrganizationsPortfolioPortfolioIDDeleteParams) (api.V1OrganizationsPortfolioPortfolioIDDeleteRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsPortfolioPortfolioIDPut(ctx context.Context, req *api.V1OrganizationsPortfolioPortfolioIDPutReq, params api.V1OrganizationsPortfolioPortfolioIDPutParams) (api.V1OrganizationsPortfolioPortfolioIDPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsPortfolioPost(ctx context.Context, req *api.V1OrganizationsPortfolioPostReq) (api.V1OrganizationsPortfolioPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDPortfolioGet(ctx context.Context, params api.V1OrganizationsOrganizationIDPortfolioGetParams) (api.V1OrganizationsOrganizationIDPortfolioGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}