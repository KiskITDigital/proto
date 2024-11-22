package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDCommentsPost(
	ctx context.Context,
	req *api.V1TendersTenderIDCommentsPostReq,
	params api.V1TendersTenderIDCommentsPostParams,
) (api.V1TendersTenderIDCommentsPostRes, error) {

	err := h.svc.CreateComment(ctx, service.CommentCreateParams{
		TenderID:    params.TenderID,
		Content:     req.Content,
		Attachments: req.Attachments,
	})
	if err != nil {
		return nil, fmt.Errorf("comment: %w", err)
	}

	return &api.V1TendersTenderIDCommentsPostOK{}, nil
}
