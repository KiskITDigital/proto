package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1TendersTenderIDCommentsGet(ctx context.Context, params api.V1TendersTenderIDCommentsGetParams) (api.V1TendersTenderIDCommentsGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1TendersTenderIDCommentsPost(ctx context.Context, req *api.V1TendersTenderIDCommentsPostReq, params api.V1TendersTenderIDCommentsPostParams) (api.V1TendersTenderIDCommentsPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1TendersTenderIDRespondPost(ctx context.Context, req *api.V1TendersTenderIDRespondPostReq, params api.V1TendersTenderIDRespondPostParams) (api.V1TendersTenderIDRespondPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1TendersVerificationsGet(ctx context.Context, params api.V1TendersVerificationsGetParams) (api.V1TendersVerificationsGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
