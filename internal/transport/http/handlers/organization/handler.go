package organization

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger              *slog.Logger
	organizationService OrganizationService
	verificationService VerificationService
}

type OrganizationService interface {
	Get(ctx context.Context, params service.OrganizationGetParams) ([]models.Organization, error)
	GetByID(ctx context.Context, id int) (models.Organization, error)
	UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error
	UpdateContacts(ctx context.Context, params service.OrganizationUpdateContactsParams) error
	CreateVerificationRequest(ctx context.Context, params service.OrganizationCreateVerificationRequestParams) error
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
