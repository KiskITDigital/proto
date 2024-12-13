package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	tenderStore       TenderStore
	commentStore      CommentStore
	verificationStore VerificationStore
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
	Count(ctx context.Context, qe store.QueryExecutor, params store.TenderGetCountParams) (int, error)
}

type CommentStore interface {
	CreateComment(ctx context.Context, qe store.QueryExecutor, params store.CommentCreateParams) error
	GetComments(ctx context.Context, qe store.QueryExecutor, params store.CommentGetParams) ([]models.Comment, error)
}

type VerificationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error
}

func New(
	psql DBTX,
	tenderStore TenderStore,
	commentStore CommentStore,
	verificationStore VerificationStore,
) *Service {
	return &Service{
		psql:              psql,
		tenderStore:       tenderStore,
		commentStore:      commentStore,
		verificationStore: verificationStore,
	}
}
