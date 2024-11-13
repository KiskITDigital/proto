package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error) {
	id, err := s.tenderStore.Create(ctx, s.psql.DB(), store.TenderCreateParams{
		Name:            params.Name,
		CityID:          params.CityID,
		OrganizationID:  params.OrganizationID,
		ServiceIDs:      params.ServiceIDs,
		ObjectIDs:       params.ObjectIDs,
		Price:           params.Price,
		IsContractPrice: params.IsContractPrice,
		IsNDSPrice:      params.IsNDSPrice,
		IsDraft:         params.IsDraft,
		FloorSpace:      params.FloorSpace,
		Description:     params.Description,
		Wishes:          params.Wishes,
		Specification:   params.Specification,
		Attachments:     params.Attachments,
		ReceptionStart:  params.ReceptionStart,
		ReceptionEnd:    params.ReceptionEnd,
		WorkStart:       params.WorkStart,
		WorkEnd:         params.WorkEnd,
	})
	if err != nil {
		return models.Tender{}, fmt.Errorf("create tender: %w", err)
	}

	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), id)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get tender: %w", err)
	}

	return tender, nil
}
