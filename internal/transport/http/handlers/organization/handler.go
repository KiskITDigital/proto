package organization

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	organizationService "gitlab.ubrato.ru/ubrato/core/internal/service/organization"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	Get(ctx context.Context, params organizationService.OrganizationGetParams) ([]models.Organization, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
