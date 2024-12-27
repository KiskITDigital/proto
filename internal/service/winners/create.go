package winners

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.WinnersCreateParams) (models.Winners, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Winners{}, cerr.Wrap(err, cerr.CodeNotFound, "tender not found", nil)
		}
		return models.Winners{}, err
	}

	if tender.Organization.ID != contextor.GetOrganizationID(ctx) {
		return models.Winners{}, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to add the winner", nil)
	}

	count, err := s.winnersStore.Count(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		return models.Winners{}, fmt.Errorf("failed to count winners: %w", err)
	}

	if count >= 3 {
		return models.Winners{}, cerr.Wrap(
			errors.New("winners limit reached"), cerr.CodeUnprocessableEntity, "Превышен лимит победителей.", nil)
	}

	var createdWinner models.Winners

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		winner, err := s.winnersStore.Create(ctx, qe, store.WinnersCreateParams{
			TenderID:       params.TenderID,
			OrganizationID: params.OrganizationID,
		})
		if err != nil {
			return fmt.Errorf("winner create: %w", err)
		}

		createdWinner = winner

		err = s.respondStore.Update(ctx, qe, store.RespondUpdateParams{
			IsWinner:       true,
			TenderID:       params.TenderID,
			OrganizationID: params.OrganizationID,
		})
		if err != nil {
			return fmt.Errorf("update is_winner status in respond table: %w", err)
		}

		return nil
	})
	if err != nil {
		return models.Winners{}, fmt.Errorf("create winner: %w", err)
	}

	return createdWinner, nil
}
