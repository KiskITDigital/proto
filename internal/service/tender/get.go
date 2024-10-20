package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetByID(ctx context.Context, tenderID int) (models.Tender, error) {
	return s.tenderStore.GetByID(ctx, s.psql.DB(), tenderID)
}
