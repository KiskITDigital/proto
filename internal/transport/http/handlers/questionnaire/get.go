package questionnaire

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1QuestionnaireGet(ctx context.Context, params api.V1QuestionnaireGetParams) (api.V1QuestionnaireGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	questionnaires, err := h.questionnaireService.Get(ctx, service.QuestionnaireGetParams{
		Offset: uint64(params.Offset.Or(0)),
		Limit:  uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get questionnaires: %w", err)
	}

	return &api.V1QuestionnaireGetOK{
		Data: convert.Slice[[]models.Questionnaire, []api.Questionnaire](questionnaires, models.ConvertQuestionnaireToAPI),
	}, nil

}
