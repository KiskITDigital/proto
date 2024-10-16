package auth

import (
	"context"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/auth"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	userStore         UserStore
	organizationStore OrganizationStore
	sessionStore      SessionStore
	dadataGateway     DadataGateway
	tokenAuthorizer   TokenAuthorizer
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type UserStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, user models.User) (models.User, error)
}

type OrganizationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, organization models.Organization) (models.Organization, error)
}

type SessionStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, session models.Session) (models.Session, error)
}

type DadataGateway interface {
	GetOrganization(ctx context.Context, INN string) (models.Organization, error)
}

type TokenAuthorizer interface {
	GenerateToken(payload auth.Payload) (string, error)
	GetRefreshTokenDurationLifetime() time.Duration
}

func New(
	psql DBTX,
	userStore UserStore,
	organizationStore OrganizationStore,
	sessionStore SessionStore,
	dadataGateway DadataGateway,
	tokenAuthorizer TokenAuthorizer,
) *Service {
	return &Service{
		psql:              psql,
		userStore:         userStore,
		organizationStore: organizationStore,
		sessionStore:      sessionStore,
		dadataGateway:     dadataGateway,
	}
}
