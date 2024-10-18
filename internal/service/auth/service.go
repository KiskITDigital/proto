package auth

import (
	"context"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"golang.org/x/exp/rand"
)

type Service struct {
	psql              DBTX
	userStore         UserStore
	organizationStore OrganizationStore
	sessionStore      SessionStore
	dadataGateway     DadataGateway
}

const (
	RefreshTokenLifetime = 7 * 24 * time.Hour
	sessionLength        = 32
)

var sessionRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!_-")

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type UserStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error)
	GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) (models.User, error)
}

type OrganizationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, organization store.OrganizationCreateParams) (models.Organization, error)
}

type SessionStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, session store.SessionCreateParams) (models.Session, error)
}

type DadataGateway interface {
	FindByINN(ctx context.Context, INN string) (dadata.FindByInnResponse, error)
}

func New(
	psql DBTX,
	userStore UserStore,
	organizationStore OrganizationStore,
	sessionStore SessionStore,
	dadataGateway DadataGateway,
) *Service {
	return &Service{
		psql:              psql,
		userStore:         userStore,
		organizationStore: organizationStore,
		sessionStore:      sessionStore,
		dadataGateway:     dadataGateway,
	}
}

func randSessionID(n int) string {
	rand.Seed(uint64(time.Now().Unix()))
	b := make([]rune, n)
	for i := range b {
		b[i] = sessionRunes[rand.Intn(len(sessionRunes))]
	}
	return string(b)
}
