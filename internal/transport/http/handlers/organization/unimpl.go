package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

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