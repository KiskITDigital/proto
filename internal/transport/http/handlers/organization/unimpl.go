package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

// Профиль
func (h *Handler) V1OrganizationsOrganizationIDProfileContractorPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContractorPutReq, params api.V1OrganizationsOrganizationIDProfileContractorPutParams) (api.V1OrganizationsOrganizationIDProfileContractorPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

// Портфолио
func (h *Handler) V1OrganizationsPortfolioPortfolioIDDelete(ctx context.Context, params api.V1OrganizationsPortfolioPortfolioIDDeleteParams) (api.V1OrganizationsPortfolioPortfolioIDDeleteRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsPortfolioPortfolioIDPut(ctx context.Context, req *api.V1OrganizationsPortfolioPortfolioIDPutReq, params api.V1OrganizationsPortfolioPortfolioIDPutParams) (api.V1OrganizationsPortfolioPortfolioIDPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDPortfolioPost(ctx context.Context, req *api.V1OrganizationsOrganizationIDPortfolioPostReq, params api.V1OrganizationsOrganizationIDPortfolioPostParams) (api.V1OrganizationsOrganizationIDPortfolioPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDPortfolioGet(ctx context.Context, params api.V1OrganizationsOrganizationIDPortfolioGetParams) (api.V1OrganizationsOrganizationIDPortfolioGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
