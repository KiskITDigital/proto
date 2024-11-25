package organization

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	organizationService "gitlab.ubrato.ru/ubrato/core/internal/service/organization"
)

type Handler struct {
	logger              *slog.Logger
	organizationService OrganizationService
	verificationService VerificationService
}

type OrganizationService interface {
	Get(ctx context.Context, params organizationService.OrganizationGetParams) ([]models.Organization, error)
	GetByID(ctx context.Context, id int) (models.Organization, error)
}

type VerificationService interface {
	Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
}

func New(
	logger *slog.Logger,
	organizationService OrganizationService,
	verificationService VerificationService) *Handler {
	return &Handler{
		logger:              logger,
		organizationService: organizationService,
		verificationService: verificationService,
	}
}
