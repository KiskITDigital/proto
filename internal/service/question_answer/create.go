package questionanswer

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.CreateQuestionAnswerParams) (models.QuestionAnswer, error) {
	if params.Type == models.QuestionAnswerTypeAnswer && !params.ParentID.Set {
		return models.QuestionAnswer{}, fmt.Errorf("parent_id must be provided for answer")
	} else if params.Type == models.QuestionAnswerTypeQuestion && params.ParentID.Set {
		return models.QuestionAnswer{}, fmt.Errorf("parent_id must not be provided for a question")
	}

	return s.questionAnswerStore.Create(ctx, s.psql.DB(), store.CreateQuestionAnswerParams{
		TenderID:             params.TenderID,
		AuthorOrganizationID: params.AuthorOrganizationID,
		ParentID:             params.ParentID,
		Type:                 params.Type,
		Content:              params.Content,
	})
}
