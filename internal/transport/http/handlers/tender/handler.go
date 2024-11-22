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
	List(ctx context.Context, params service.TenderListParams) ([]models.Tender, error)
	Respond(ctx context.Context, params service.TenderRespondParams) error
	CreateComment(ctx context.Context, params service.CommentCreateParams) error
	GetComments(ctx context.Context, params service.GetCommentParams) ([]models.Comment, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
