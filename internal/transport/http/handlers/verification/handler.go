package verification

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
	UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error
	GetByID(ctx context.Context, requestID int) (models.VerificationRequest[models.VerificationObject], error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
