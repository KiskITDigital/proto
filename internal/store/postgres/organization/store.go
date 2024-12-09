package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type (
	OrganizationStore struct {
		catalogStore CatalogStore
	}

	CatalogStore interface {
		GetCitiesByIDs(ctx context.Context, qe store.QueryExecutor, cityIDs []int) (map[int]models.City, error)
		GetObjectsByIDs(ctx context.Context, qe store.QueryExecutor, objectIDs []int) (map[int]models.Object, error)
		GetServicesByIDs(ctx context.Context, qe store.QueryExecutor, serviceIDs []int) (map[int]models.Service, error)
	}
)

func NewOrganizationStore(catalogStore CatalogStore) *OrganizationStore {
	return &OrganizationStore{catalogStore: catalogStore}
}
