package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type CatalogService struct {
	ID       int
	ParentID int
	Name     string
}

func ConvertModelCatalogServiceToApi(c CatalogService) api.Service {
	return api.Service{
		ID:       c.ID,
		ParentID: api.OptInt{Value: c.ParentID, Set: c.ParentID != 0},
		Name:     c.Name,
	}
}

type CatalogServices []CatalogService

type CatalogObject struct {
	ID       int
	ParentID int
	Name     string
}

type CatalogObjects []CatalogObject

func ConvertModelCatalogObjectToApi(c CatalogObject) api.Object {
	return api.Object{
		ID:       c.ID,
		ParentID: api.OptInt{Value: c.ParentID, Set: c.ParentID != 0},
		Name:     c.Name,
	}
}
