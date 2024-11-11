package suggest

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
)

type Service struct {
	dadataGateway DadataGateway
}

type DadataGateway interface {
	FindByINN(ctx context.Context, INN string) (dadata.FindByInnResponse, error)
}

func New(
	dadataGateway DadataGateway,

) *Service {
	return &Service{
		dadataGateway: dadataGateway,
	}
}
