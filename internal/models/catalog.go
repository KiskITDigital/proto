package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type CatalogService struct {
	ID       int
	ParentID int
	Name     string
}

func ConvertModelCatalogServiceToApi(c CatalogService) api.ServicesItem {
	return api.ServicesItem{
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

func ConvertModelCatalogObjectToApi(c CatalogObject) api.ObjectsItem {
	return api.ObjectsItem{
		ID:       c.ID,
		ParentID: api.OptInt{Value: c.ParentID, Set: c.ParentID != 0},
		Name:     c.Name,
	}
}
