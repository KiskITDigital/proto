package tender

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1TendersTenderIDQuestionAnswerPost(ctx context.Context, req *api.V1TendersTenderIDQuestionAnswerPostReq, params api.V1TendersTenderIDQuestionAnswerPostParams) (api.V1TendersTenderIDQuestionAnswerPostRes, error) {
	questionAnswer, err := h.tenderService.CreateQuestionAnswer(ctx, service.CreateQuestionAnswerParams{
		TenderID:             params.TenderID,
		AuthorOrganizationID: contextor.GetOrganizationID(ctx),
		ParentID:             models.Optional[int]{Value: req.GetParentID().Value, Set: req.GetParentID().Set},
		Type:                 models.APIToQuestionAnswerType(params.Type),
		Content:              req.Content})
	if err != nil {
		if errors.Is(err, errstore.ErrQuestionAnswerUniqueViolation) {
			return nil, cerr.Wrap(err, cerr.CodeConflict, "Ответ на вопрос уже существует", nil)
		}

		return nil, fmt.Errorf("create question-answer: %w", err)
	}

	return &api.V1TendersTenderIDQuestionAnswerPostCreated{
		Data: models.ConvertQuestionAnswerToAPI(questionAnswer),
	}, nil
}
