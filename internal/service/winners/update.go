package winners

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.WinnerUpdateParams) error {
	organizationID, err := s.winnersStore.GetOrganizationIDByWinnerID(ctx, s.psql.DB(), params.WinnerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cerr.Wrap(err, cerr.CodeNotFound, "winner not found", nil)
		}
		return err
	}

	if organizationID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to update the winner accepted status", nil)
	}

	return s.winnersStore.UpdateStatus(ctx, s.psql.DB(), store.WinnerUpdateParams{
		WinnerID: params.WinnerID,
		Accepted: params.Accepted,
	})
}
