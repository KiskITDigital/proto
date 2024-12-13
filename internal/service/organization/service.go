package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	organizationStore OrganizationStore
	verificationStore VerificationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type OrganizationStore interface {
	Get(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetParams) ([]models.Organization, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
	GetCustomer(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error)
	GetContractor(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateParams) error
}

type VerificationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error
}

func New(
	psql DBTX,
	organizationStore OrganizationStore,
	verificationStore VerificationStore) *Service {
	return &Service{
		psql:              psql,
		organizationStore: organizationStore,
		verificationStore: verificationStore,
	}
}
