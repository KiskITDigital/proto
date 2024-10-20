package tender

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	tenderService "gitlab.ubrato.ru/ubrato/core/internal/service/tender"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	Create(ctx context.Context, params tenderService.CreateParams) (models.Tender, error)
	Update(ctx context.Context, params tenderService.UpdateParams) (models.Tender, error)
	GetByID(ctx context.Context, tenderID int) (models.Tender, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
