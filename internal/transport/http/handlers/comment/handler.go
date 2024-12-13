package comment

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type (
	Handler struct {
		logger              *slog.Logger
		commentService      CommentService
		verificationService VerificationService
	}

	CommentService interface {
	}

	VerificationService interface {
		Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) (models.VerificationRequestPagination[models.VerificationObject], error)
	}
)

func New(
	logger *slog.Logger,
	commentService CommentService,
	verificationService VerificationService) *Handler {
	return &Handler{
		logger:              logger,
		commentService:      commentService,
		verificationService: verificationService,
	}
}
