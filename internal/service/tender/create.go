package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error) {
	var (
		tender models.Tender
		err    error
	)

	organizationID := contextor.GetOrganizationID(ctx)

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		tender, err = s.tenderStore.Create(ctx, qe, store.TenderCreateParams{
			Name:            params.Name,
			CityID:          params.CityID,
			OrganizationID:  organizationID,
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
			return fmt.Errorf("create tender: %w", err)
		}

		err = s.tenderStore.AppendTenderServies(ctx, qe, store.TenderServicesCreateParams{
			TenderID:    tender.ID,
			ServicesIDs: params.ServiceIDs,
		})
		if err != nil {
			return fmt.Errorf("append services to tender: %w", err)
		}

		err = s.tenderStore.AppendTenderObjects(ctx, qe, store.TenderObjectsCreateParams{
			TenderID:   tender.ID,
			ObjectsIDs: params.ObjectIDs,
		})
		if err != nil {
			return fmt.Errorf("append objects to tender: %w", err)
		}

		tender, err = s.tenderStore.GetByID(ctx, qe, tender.ID)
		if err != nil {
			return fmt.Errorf("get tender: %w", err)
		}

		return nil
	})

	if err != nil {
		return models.Tender{}, fmt.Errorf("run transaction: %w", err)
	}

	return tender, nil
}
