package catalog

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	catalogStore CatalogStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type CatalogStore interface {
	GetServices(ctx context.Context, qe store.QueryExecutor) (models.CatalogServices, error)
	GetObjects(ctx context.Context, qe store.QueryExecutor) (models.CatalogObjects, error)
}

func New(
	psql DBTX,
	catalogStore CatalogStore,
) *Service {
	return &Service{
		psql:         psql,
		catalogStore: catalogStore,
	}
}
