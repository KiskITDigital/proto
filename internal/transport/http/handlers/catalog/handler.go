package catalog

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	GetObjects(ctx context.Context) (models.CatalogObjects, error)
	GetServices(ctx context.Context) (models.CatalogServices, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
