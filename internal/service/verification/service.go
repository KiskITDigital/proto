package verification

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	verificationStore VerificationStore
	tenderStore       TenderStore
	organizationStore OrganizationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type VerificationStore interface {
	UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestUpdateStatusParams) (store.VerificationObjectUpdateStatusResult, error)
	GetOrganizationRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
	GetTendersRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
}

type TenderStore interface {
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error) 
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateVerifStatusParams) error
}

type OrganizationStore interface {
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateVerifStatusParams) error
}

func New(
	psql DBTX,
	verificationStore VerificationStore,
	tenderStore TenderStore,
	organiOrganizationStore OrganizationStore,
) *Service {
	return &Service{
		psql:              psql,
		verificationStore: verificationStore,
		tenderStore:       tenderStore,
		organizationStore: organiOrganizationStore,
	}
}
