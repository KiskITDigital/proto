package user

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
	GetByID(ctx context.Context, tenderID int) (models.RegularUser, error)
	Get(ctx context.Context) ([]models.RegularUser, error)
	ReqEmailVerification(ctx context.Context, email string) error
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
