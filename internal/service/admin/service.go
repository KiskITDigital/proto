package admin

import (
	"context"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"golang.org/x/exp/rand"
)

type Service struct {
	psql            DBTX
	adminStore      AdminStore
	tokenAuthorizer TokenAuthorizer
}

const (
	sessionLength = 32
)

var sessionRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!_-")

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type AdminStore interface {
	GetUser(ctx context.Context, qe store.QueryExecutor, params store.AdminGetUserParams) (models.AdminUser, error)
	ListUsers(ctx context.Context, qe store.QueryExecutor) ([]models.AdminUser, error)
	CreateUser(ctx context.Context, qe store.QueryExecutor, params store.AdminCreateUserParams) (models.AdminUser, error)
	GetSession(ctx context.Context, qe store.QueryExecutor, sessionID string) (models.Session, error)
	CreateSession(ctx context.Context, qe store.QueryExecutor, params store.SessionCreateParams) (models.Session, error)
	UpdateSession(ctx context.Context, qe store.QueryExecutor, params store.SessionUpdateParams) (models.Session, error)
}

type TokenAuthorizer interface {
	GenerateToken(payload token.Payload) (string, error)
	GetRefreshTokenDurationLifetime() time.Duration
	ValidateToken(rawToken string) (token.Claims, error)
}

func New(
	psql DBTX,
	adminStore AdminStore,
	tokenAuthorizer TokenAuthorizer,
) *Service {
	return &Service{
		psql:            psql,
		adminStore:      adminStore,
		tokenAuthorizer: tokenAuthorizer,
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
