package winners

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	winnersStore WinnersStore
	tenderStore  TenderStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
}

type WinnersStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.WinnersCreateParams) (models.Winners, error)
	Get(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.Winners, error)
	UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.WinnerUpdateParams) error
	GetOrganizationIDByWinnerID(ctx context.Context, qe store.QueryExecutor, winnerID int) (int, error)
	Count(ctx context.Context, qe store.QueryExecutor, tenderID int) (int, error)
}

type TenderStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error)
}

func New(
	psql DBTX,
	winnersStore WinnersStore,
	tenderStore TenderStore,
) *Service {
	return &Service{
		psql:         psql,
		winnersStore: winnersStore,
		tenderStore:  tenderStore,
	}
}
