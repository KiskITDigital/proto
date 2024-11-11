package suggest

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (h *Handler) V1SuggestCompanyGet(ctx context.Context, params api.V1SuggestCompanyGetParams) (api.V1SuggestCompanyGetRes, error) {
	companies, err := h.svc.SuggestByINN(ctx, string(params.Inn))
	if err != nil {
		return nil, fmt.Errorf("search by inn: %w", err)
	}

	if len(companies.Suggestions) == 0 {
		return nil, cerr.Wrap(
			fmt.Errorf("search by inn"),
			cerr.CodeNotFound,
			fmt.Sprintf("company with inn %s not found", params.Inn),
			nil,
		)
	}

	return &api.V1SuggestCompanyGetOK{
		Data: api.V1SuggestCompanyGetOKData{
			Name: companies.Suggestions[0].Data.Name.ShortWithOpf,
		},
	}, nil
}
