package tender

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger                *slog.Logger
	tenderService         TenderService
	verificationService   VerificationService
	questionAnswerService QuestionAnswerService
	respondService        RespondService
}

type TenderService interface {
	Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error)
	Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error)
	GetByID(ctx context.Context, tenderID int) (models.Tender, error)
	List(ctx context.Context, params service.TenderListParams) (models.TendersPagination, error)
	CreateComment(ctx context.Context, params service.CommentCreateParams) error
	GetComments(ctx context.Context, params service.GetCommentParams) ([]models.Comment, error)
}

type QuestionAnswerService interface {
	Create(ctx context.Context, params service.CreateQuestionAnswerParams) (models.QuestionAnswer, error)
	Get(ctx context.Context, tenderID int) ([]models.QuestionWithAnswer, error)
}

type VerificationService interface {
	Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) (models.VerificationRequestPagination[models.VerificationObject], error)
}

type RespondService interface {
	Create(ctx context.Context, params service.RespondCreateParams) error
	Get(ctx context.Context, params service.RespondGetParams) (models.RespondPagination, error)
}

func New(
	logger *slog.Logger,
	tenderService TenderService,
	verificationService VerificationService,
	questionAnswerService QuestionAnswerService,
	respondService RespondService,
) *Handler {
	return &Handler{
		logger:                logger,
		tenderService:         tenderService,
		verificationService:   verificationService,
		questionAnswerService: questionAnswerService,
		respondService:        respondService,
	}
}
