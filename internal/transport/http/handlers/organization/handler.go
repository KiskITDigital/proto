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
	portfolioService    PortfolioService
}

type OrganizationService interface {
	Get(ctx context.Context, params service.OrganizationGetParams) (models.OrganizationsPagination, error)
	GetByID(ctx context.Context, id int) (models.Organization, error)
	GetCustomer(ctx context.Context, organizationId int) (models.Organization, error)
	GetContractor(ctx context.Context, organizationId int) (models.Organization, error)
	UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error
	UpdateContacts(ctx context.Context, params service.OrganizationUpdateContactsParams) error
	CreateVerificationRequest(ctx context.Context, params service.OrganizationCreateVerificationRequestParams) error
	UpdateCustomer(ctx context.Context, params service.OrganizationUpdateCustomerParams) (models.Organization, error)
	UpdateContractor(ctx context.Context, params service.OrganizationUpdateContractorParams) (models.Organization, error)
}

type VerificationService interface {
	Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) (models.VerificationRequestPagination[models.VerificationObject], error)
}

type PortfolioService interface {
	Create(ctx context.Context, params service.PortfolioCreateParams) (models.Portfolio, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, params service.PortfolioGetParams) ([]models.Portfolio, error)
	Update(ctx context.Context, params service.PortfolioUpdateParams) (models.Portfolio, error)
}

func New(
	logger *slog.Logger,
	organizationService OrganizationService,
	verificationService VerificationService,
	portfolioService PortfolioService) *Handler {
	return &Handler{
		logger:              logger,
		organizationService: organizationService,
		verificationService: verificationService,
		portfolioService:    portfolioService,
	}
}
