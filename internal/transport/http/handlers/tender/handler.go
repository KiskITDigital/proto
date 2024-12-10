package tender

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger *slog.Logger
	tenderService TenderService
	verificationService VerificationService
}

type TenderService interface {
	Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error)
	Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error)
	GetByID(ctx context.Context, tenderID int) (models.Tender, error)
	List(ctx context.Context, params service.TenderListParams) (models.TendersRes, error) 
	Respond(ctx context.Context, params service.TenderRespondParams) error
	CreateComment(ctx context.Context, params service.CommentCreateParams) error
	GetComments(ctx context.Context, params service.GetCommentParams) ([]models.Comment, error)
}

type VerificationService interface {
	Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error)
}

func New(
	logger *slog.Logger, 
	tenderService TenderService, 
	verificationService VerificationService) *Handler {
	return &Handler{
		logger: logger,
		tenderService: tenderService,
		verificationService: verificationService,
	}
}
