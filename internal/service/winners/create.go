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

	return s.winnersStore.Create(ctx, s.psql.DB(), store.WinnersCreateParams{
		TenderID:       params.TenderID,
		OrganizationID: params.OrganizationID,
	})
}
