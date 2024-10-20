package tender

import (
	"context"
	"fmt"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type CreateParams struct {
	Name            string
	CityID          int
	OrganizationID  int
	Price           float64
	IsContractPrice bool
	IsNDSPrice      bool
	FloorSpace      int
	Description     string
	Wishes          string
	Specification   string
	Attachments     []string
	ServiceIDs      []int
	ObjectIDs       []int
	ReceptionStart  time.Time
	ReceptionEnd    time.Time
	WorkStart       time.Time
	WorkEnd         time.Time
}

func (s *Service) Create(ctx context.Context, params CreateParams) (models.Tender, error) {
	var (
		tender models.Tender
		err    error
	)

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		tender, err = s.tenderStore.Create(ctx, qe, store.TenderCreateParams{
			Name:            params.Name,
			CityID:          params.CityID,
			OrganizationID:  params.OrganizationID,
			Price:           params.Price,
			IsContractPrice: params.IsContractPrice,
			IsNDSPrice:      params.IsNDSPrice,
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
