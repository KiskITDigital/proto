package questionanswer

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) Get(ctx context.Context, tenderID int) ([]models.QuestionWithAnswer, error) {
	return s.questionAnswerStore.Get(ctx, s.psql.DB(), tenderID)
}
