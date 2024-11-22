package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	tenderStore  TenderStore
	commentStore CommentStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type TenderStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (int, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, tenderID int) (models.Tender, error)
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (int, error)
	CreateResponse(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateResponseParams) error
	CreateComment(ctx context.Context, qe store.QueryExecutor, params store.CommentCreateParams) error
}

type CommentStore interface {
	CreateComment(ctx context.Context, qe store.QueryExecutor, params store.CommentCreateParams) error
}

func New(
	psql DBTX,
	tenderStore TenderStore,
	commentStore CommentStore,
) *Service {
	return &Service{
		psql:         psql,
		tenderStore:  tenderStore,
		commentStore: commentStore,
	}
}
