package suggest

import (
	"context"
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	Company(ctx context.Context, INN string) (string, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
