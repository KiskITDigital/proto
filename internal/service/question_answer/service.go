package questionanswer

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql                DBTX
	questionAnswerStore QuestionAnswerStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type QuestionAnswerStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.CreateQuestionAnswerParams) (models.QuestionAnswer, error)
	Get(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.QuestionWithAnswer, error)
}

func New(
	psql DBTX,
	questionAnswerStore QuestionAnswerStore,
) *Service {
	return &Service{
		psql:                psql,
		questionAnswerStore: questionAnswerStore,
	}
}
