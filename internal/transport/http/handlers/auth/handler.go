package auth

import (
	"context"

	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

type Handler struct {
	svc Service
}

type Service interface {
	SignUp(ctx context.Context, params authService.SignUpParams) (authService.SignUpResult, error)
}

func New(svc Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
