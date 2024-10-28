package tender

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error)
	Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error)
	GetByID(ctx context.Context, tenderID int) (models.Tender, error)
	Get(ctx context.Context, params service.TenderGetParams) ([]models.Tender, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
