package auth

import (
	"context"
	"log/slog"

	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	SignUp(ctx context.Context, params authService.SignUpParams) (authService.SignUpResult, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
