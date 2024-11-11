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
	Get(ctx context.Context, qe store.QueryExecutor, params store.TenderGetParams) ([]models.Tender, error)
	GetTendersServices(ctx context.Context, qe store.QueryExecutor, tenderIDs []int) ([]models.TenderService, error)
	GetTendersObjects(ctx context.Context, qe store.QueryExecutor, tenderIDs []int) ([]models.TenderObject, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (models.Tender, error)
	DeleteTenderObjects(ctx context.Context, qe store.QueryExecutor, params store.TenderObjectsDeleteParams) error
	DeleteTenderServices(ctx context.Context, qe store.QueryExecutor, params store.TenderServicesDeleteParams) error
	CreateResponse(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateResponseParams) error
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
