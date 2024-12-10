package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1TendersTenderIDQuestionAnswerGet(ctx context.Context, params api.V1TendersTenderIDQuestionAnswerGetParams) (api.V1TendersTenderIDQuestionAnswerGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
