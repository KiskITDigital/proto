package questionanswer

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.CreateQuestionAnswerParams) (models.QuestionAnswer, error) {
	if params.Type == models.QuestionAnswerTypeAnswer {
		question, err := s.questionAnswerStore.GetByID(ctx, s.psql.DB(), params.ParentID.Value)
		if err != nil {
			return models.QuestionAnswer{}, fmt.Errorf("get question: %w", err)
		}

		if question.Question.VerificationStatus != models.VerificationStatusApproved {
			return models.QuestionAnswer{}, cerr.Wrap(
				fmt.Errorf("question status not approved"), 
				cerr.CodeUnprocessableEntity, 
				"Нельзя отправить ответ, так как вопрос не прошел модерацию", nil)
		}
	}

	var questionAnswer models.QuestionAnswer
	if err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		newQuestionAnswer, err := s.questionAnswerStore.Create(ctx, qe, store.CreateQuestionAnswerParams{
			TenderID:             params.TenderID,
			AuthorOrganizationID: params.AuthorOrganizationID,
			ParentID:             params.ParentID,
			Type:                 params.Type,
			Content:              params.Content})
		if err != nil {
			return fmt.Errorf("create question-answer %w", err)
		}
		questionAnswer = newQuestionAnswer

		err = s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:   newQuestionAnswer.ID,
			ObjectType: models.ObjectTypeQuestionAnswer})
		if err != nil {
			return fmt.Errorf("create verification request: %w", err)
		}

		return nil
	}); err != nil {
		return models.QuestionAnswer{}, fmt.Errorf("run transaction: %w", err)
	}

	return questionAnswer, nil
}
