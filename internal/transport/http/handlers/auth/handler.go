package auth

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	SignUp(ctx context.Context, params authService.SignUpParams) (authService.SignUpResult, error)
	SignIn(ctx context.Context, params authService.SignInParams) (authService.SignInResult, error)
	Refresh(ctx context.Context, sessionToken string) (authService.SignInResult, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
