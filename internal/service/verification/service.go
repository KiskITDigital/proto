package verification

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql                DBTX
	verificationStore   VerificationStore
	tenderStore         TenderStore
	additionStore       AdditionStore
	organizationStore   OrganizationStore
	questionAnswerStore QuestionAnswerStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type VerificationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error
	UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestUpdateStatusParams) (store.VerificationObjectUpdateStatusResult, error)
	GetWithEmptyObject(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
	GetByIDWithEmptyObject(ctx context.Context, qe store.QueryExecutor, requestID int) (models.VerificationRequest[models.VerificationObject], error)
	Count(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetCountParams) (int, error)
}

type TenderStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error)
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error)
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateVerifStatusParams) error
}

type AdditionStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Addition, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.AdditionGetParams) ([]models.Addition, error)
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.AdditionUpdateVerifStatusParams) error
}

type OrganizationStore interface {
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateVerifStatusParams) error
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetParams) ([]models.Organization, error)
}

type QuestionAnswerStore interface {
	Get(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerGetParams) ([]models.QuestionWithAnswer, error)
	UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerVerifStatusUpdateParams) error 
}

func New(
	psql DBTX,
	verificationStore VerificationStore,
	tenderStore TenderStore,
	additionStore AdditionStore,
	organiOrganizationStore OrganizationStore,
	questionAnswerStore QuestionAnswerStore,
) *Service {
	return &Service{
		psql:                psql,
		verificationStore:   verificationStore,
		tenderStore:         tenderStore,
		additionStore:       additionStore,
		organizationStore:   organiOrganizationStore,
		questionAnswerStore: questionAnswerStore,
	}
}
