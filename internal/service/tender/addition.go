package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) CreateAddition(ctx context.Context, params service.AdditionCreateParams) error {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		return fmt.Errorf("get tender: %w", err)
	}

	if tender.Organization.ID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "Недостаточно прав для добавления дополнительной информации", nil)
	}

	if err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		additionID, err := s.additionStore.CreateAddition(ctx, qe, store.AdditionCreateParams{
			TenderID:       params.TenderID,
			Title:          params.Title,
			Content:        params.Content,
			Attachments:    params.Attachments})
		if err != nil {
			return fmt.Errorf("creating addition %w", err)
		}

		err = s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:   additionID,
			ObjectType: models.ObjectTypeAddition})
		if err != nil {
			return fmt.Errorf("create verification request: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("run transaction: %w", err)
	}

	return nil
}

func (s *Service) GetAdditions(ctx context.Context, params service.GetAdditionParams) ([]models.Addition, error) {
	additions, err := s.additionStore.Get(ctx, s.psql.DB(), store.AdditionGetParams{
		TenderID:     models.NewOptional(params.TenderID),
		VerifiedOnly: true})
	if err != nil {
		return nil, fmt.Errorf("get Addition: %w", err)
	}

	return additions, nil
}
