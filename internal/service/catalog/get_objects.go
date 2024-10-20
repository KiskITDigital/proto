package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetObjects(ctx context.Context) (models.CatalogObjects, error) {
	objects, err := s.catalogStore.GetObjects(ctx, s.psql.DB())
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return objects, nil
}
