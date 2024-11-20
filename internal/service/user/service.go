package user

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql      DBTX
	userStore UserStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type UserStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error)
	GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.RegularUser, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.FullUser, error)
}

func New(
	psql DBTX,
	userStore UserStore,
) *Service {
	return &Service{
		psql:      psql,
		userStore: userStore,
	}
}
