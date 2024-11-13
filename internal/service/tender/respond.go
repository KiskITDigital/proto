package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Respond(ctx context.Context, params service.TenderRespondParams) error {
	err := s.tenderStore.CreateResponse(ctx, s.psql.DB(), store.TenderCreateResponseParams{
		TenderID:       params.TenderID,
		OrganizationID: params.OrganizationID,
		Price:          params.Price,
		IsNds:          params.IsNds,
	})
	if err != nil {
		return fmt.Errorf("create response: %w", err)
	}

	return nil
}
