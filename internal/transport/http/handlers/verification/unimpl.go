package verification

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1VerificationsRequestIDAprovePost(ctx context.Context, params api.V1VerificationsRequestIDAprovePostParams) (api.V1VerificationsRequestIDAprovePostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1VerificationsRequestIDDenyPost(ctx context.Context, params api.V1VerificationsRequestIDDenyPostParams) (api.V1VerificationsRequestIDDenyPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
