package suggest

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
)

func (s *Service) SuggestByINN(ctx context.Context, inn string) (dadata.FindByInnResponse, error) {
	return s.dadataGateway.FindByINN(ctx, inn)
}
