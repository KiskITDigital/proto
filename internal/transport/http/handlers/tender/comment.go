package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDCommentsPost(
	ctx context.Context,
	req *api.V1TendersTenderIDCommentsPostReq,
	params api.V1TendersTenderIDCommentsPostParams,
) (api.V1TendersTenderIDCommentsPostRes, error) {

	err := h.tenderService.CreateComment(ctx, service.CommentCreateParams{
		TenderID:    params.TenderID,
		Content:     req.Content,
		Attachments: req.Attachments,
	})
	if err != nil {
		return nil, fmt.Errorf("comment: %w", err)
	}

	return &api.V1TendersTenderIDCommentsPostOK{}, nil
}

func (h *Handler) V1TendersTenderIDCommentsGet(
	ctx context.Context,
	params api.V1TendersTenderIDCommentsGetParams,
) (api.V1TendersTenderIDCommentsGetRes, error) {
	comment, err := h.tenderService.GetComments(ctx, service.GetCommentParams{TenderID: params.TenderID})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersTenderIDCommentsGetOK{
		Data: convert.Slice[[]models.Comment, []api.Comment](comment, models.ConvertCommentModelToApi),
	}, nil
}
