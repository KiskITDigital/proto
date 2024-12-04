package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1UsersUserIDPut(ctx context.Context, req *api.V1UsersUserIDPutReq, params api.V1UsersUserIDPutParams) (api.V1UsersUserIDPutRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
