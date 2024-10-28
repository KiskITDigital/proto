package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error) {
	var (
		tender models.Tender
		err    error
	)

	organizationID := contextor.GetOrganizationID(ctx)

	tender, err = s.tenderStore.GetByID(ctx, s.psql.DB(), params.ID)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get tender: %w", err)
	}

	if tender.Organization.ID != organizationID {
		cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit this tender", nil)
	}

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		_, err := s.tenderStore.Update(ctx, qe, store.TenderUpdateParams{
			ID:              params.ID,
			Name:            params.Name,
			Price:           params.Price,
			IsContractPrice: params.IsContractPrice,
			IsNDSPrice:      params.IsNDSPrice,
			IsDraft:         params.IsDraft,
			CityID:          params.CityID,
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
			return fmt.Errorf("update tender: %w", err)
		}

		if params.ServiceIDs.Set {
			err = s.updateServices(ctx, qe, tender.ID, tender.Services)
			if err != nil {
				return fmt.Errorf("update tender services: %w", err)
			}
		}

		if params.ServiceIDs.Set {
			err = s.updateObjects(ctx, qe, tender.ID, tender.Objects)
			if err != nil {
				return fmt.Errorf("update tender services: %w", err)
			}
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

func (s *Service) updateServices(ctx context.Context, qe store.QueryExecutor, tenderID int, services []models.TenderService) error {
	err := s.tenderStore.DeleteTenderServices(ctx, qe, store.TenderServicesDeleteParams{
		TenderID:    tenderID,
		ServicesIDs: convert.Slice[[]models.TenderService, []int](services, func(ts models.TenderService) int { return ts.ID }),
	})
	if err != nil {
		return fmt.Errorf("delete services: %w", err)
	}

	err = s.tenderStore.AppendTenderServies(ctx, qe, store.TenderServicesCreateParams{
		TenderID:    tenderID,
		ServicesIDs: convert.Slice[[]models.TenderService, []int](services, func(ts models.TenderService) int { return ts.ID }),
	})
	if err != nil {
		return fmt.Errorf("append services: %w", err)
	}

	return nil
}

func (s *Service) updateObjects(ctx context.Context, qe store.QueryExecutor, tenderID int, services []models.TenderObject) error {
	err := s.tenderStore.DeleteTenderObjects(ctx, qe, store.TenderObjectsDeleteParams{
		TenderID:   tenderID,
		ObjectsIDs: convert.Slice[[]models.TenderObject, []int](services, func(to models.TenderObject) int { return to.ID }),
	})
	if err != nil {
		return fmt.Errorf("delete services: %w", err)
	}

	err = s.tenderStore.AppendTenderServies(ctx, qe, store.TenderServicesCreateParams{
		TenderID:    tenderID,
		ServicesIDs: convert.Slice[[]models.TenderObject, []int](services, func(to models.TenderObject) int { return to.ID }),
	})
	if err != nil {
		return fmt.Errorf("append objects: %w", err)
	}

	return nil
}
