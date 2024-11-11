package suggest

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	SuggestByINN(ctx context.Context, inn string) (dadata.FindByInnResponse, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
