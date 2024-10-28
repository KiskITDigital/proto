package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetByID(ctx context.Context, tenderID int) (models.Tender, error) {
	return s.tenderStore.GetByID(ctx, s.psql.DB(), tenderID)
}

func (s *Service) Get(ctx context.Context) ([]models.Tender, error) {
	return s.tenderStore.Get(ctx, s.psql.DB())
}
