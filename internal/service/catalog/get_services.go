package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetServices(ctx context.Context) (models.CatalogServices, error) {
	services, err := s.catalogStore.GetServices(ctx, s.psql.DB())
	if err != nil {
		return nil, fmt.Errorf("get services catalog: %w", err)
	}

	return services, nil
}
