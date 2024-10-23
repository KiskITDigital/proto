package survey

import (
	"log/slog"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
