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
	commentStore CommentStore
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
	GetCommentRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
	GetWithEmptyObject(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
	GetByIDWithEmptyObject(ctx context.Context, qe store.QueryExecutor, requestID int) (models.VerificationRequest[models.VerificationObject], error)
}

type TenderStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error)
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error)
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateVerifStatusParams) error
}

type CommentStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Comment, error)
}

type OrganizationStore interface {
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateVerifStatusParams) error
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
}

func New(
	psql DBTX,
	verificationStore VerificationStore,
	tenderStore TenderStore,
	commentStore CommentStore,
	organiOrganizationStore OrganizationStore,
) *Service {
	return &Service{
		psql:              psql,
		verificationStore: verificationStore,
		tenderStore:       tenderStore,
		commentStore:      commentStore,
		organizationStore: organiOrganizationStore,
	}
}
