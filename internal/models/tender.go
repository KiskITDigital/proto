package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type Tender struct {
	ID              int
	Name            string
	City            City
	Organization    Organization
	Price           float64
	IsContractPrice bool
	IsNDSPrice      bool
	FloorSpace      int
	Description     string
	Wishes          string
	Specification   string
	Attachments     []string
	Services        []TenderService
	Objects         []TenderObject
	Verified        bool
	ReceptionStart  time.Time
	ReceptionEnd    time.Time
	WorkStart       time.Time
	WorkEnd         time.Time
	CreatedAt       time.Time
}

type City struct {
	ID     int
	Name   string
	Region Region
}

func ConvertCityModelToApi(city City) api.City {
	return api.City{
		ID:   city.ID,
		Name: city.Name,
		Region: api.OptRegion{
			Value: ConvertRegionModelToApi(city.Region),
			Set:   city.Region.ID != 0,
		},
	}
}

type Region struct {
	ID   int
	Name string
}

func ConvertRegionModelToApi(region Region) api.Region {
	return api.Region{
		ID:   region.ID,
		Name: region.Name,
	}
}

type TenderObject struct {
	ID       int
	ParentID int
	Name     string
}

type TenderService struct {
	ID       int
	ParentID int
	Name     string
}

func ConvertTenderModelToApi(tender Tender) api.Tender {
	return api.Tender{
		ID:              tender.ID,
		Name:            tender.Name,
		City:            tender.City.Name,
		Organization:    api.OptOrganization{Value: ConvertOrganizationModelToApi(tender.Organization), Set: tender.Organization.ID != 0},
		Region:          tender.City.Region.Name,
		Price:           tender.Price,
		IsContractPrice: tender.IsContractPrice,
		IsNdsPrice:      tender.IsNDSPrice,
		FloorSpace:      tender.FloorSpace,
		Description:     tender.Description,
		Wishes:          tender.Wishes,
		Specification:   api.URL(tender.Specification),
		Attachments: convert.Slice[[]string, []api.URL](
			tender.Attachments, func(u string) api.URL { return api.URL(u) },
		),
		Services: convert.Slice[[]TenderService, api.Services](
			tender.Services, ConvertTenderServiceModelToApi,
		),
		Objects: convert.Slice[[]TenderObject, api.Objects](
			tender.Objects, ConvertTenderObjectModelToApi,
		),
		ReceptionStart: tender.ReceptionStart,
		ReceptionEnd:   tender.ReceptionEnd,
		WorkStart:      tender.WorkStart,
		WorkEnd:        tender.WorkEnd,
	}
}

func ConvertTenderServiceModelToApi(service TenderService) api.Service {
	return api.Service{
		ID:       service.ID,
		ParentID: api.OptInt{Value: service.ParentID, Set: service.ParentID != 0},
		Name:     service.Name,
	}
}

func ConvertTenderObjectModelToApi(object TenderObject) api.Object {
	return api.Object{
		ID:       object.ID,
		ParentID: api.OptInt{Value: object.ParentID, Set: object.ParentID != 0},
		Name:     object.Name,
	}
}
