package admin

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/service/admin"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	SignIn(ctx context.Context, params admin.SignInParams) (admin.SignInResult, error)
	Refresh(ctx context.Context, sessionID string) (admin.SignInResult, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error)
	GetUser(ctx context.Context, params admin.GetUserParams) (admin.GetUserResult, error)
	ListUsers(ctx context.Context) (admin.ListUsersResult, error)
	CreateUser(ctx context.Context, params admin.CreateUserParams) (admin.CreateUserResult, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
