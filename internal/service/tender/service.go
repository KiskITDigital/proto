package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql        DBTX
	tenderStore TenderStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type TenderStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (models.Tender, error)
	AppendTenderServies(ctx context.Context, qe store.QueryExecutor, params store.TenderServicesCreateParams) error
	AppendTenderObjects(ctx context.Context, qe store.QueryExecutor, params store.TenderObjectsCreateParams) error
	GetByID(ctx context.Context, qe store.QueryExecutor, tenderID int) (models.Tender, error)
	GetTenderServices(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.TenderService, error)
	GetTenderObjects(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.TenderObject, error)
}

func New(
	psql DBTX,
	tenderStore TenderStore,
) *Service {
	return &Service{
		psql:        psql,
		tenderStore: tenderStore,
	}
}
