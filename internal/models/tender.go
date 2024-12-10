package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type TendersRes struct {
	Tenders []Tender
	Pagination Pagination
}

type Tender struct {
	VerificationObject

	ID                 int
	Name               string
	City               City
	Organization       Organization
	WinnerOrganization Optional[Organization]
	Price              int
	IsContractPrice    bool
	IsNDSPrice         bool
	IsDraft            bool
	FloorSpace         int
	Description        string
	Wishes             string
	Specification      string
	Attachments        []string
	Services           []Service
	Objects            []Object
	VerificationStatus VerificationStatus
	Status             int
	ReceptionStart     time.Time
	ReceptionEnd       time.Time
	WorkStart          time.Time
	WorkEnd            time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (t Tender) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:   api.TenderVerificationRequestObject,
		Tender: ConvertTenderModelToApi(t),
	}
}

type Object struct {
	ID       int
	ParentID int
	Name     string
}

type Service struct {
	ID       int
	ParentID int
	Name     string
}

func ConvertTenderModelToApi(tender Tender) api.Tender {
	tenderApi := api.Tender{
		ID:              tender.ID,
		Name:            tender.Name,
		City:            ConvertCityModelToApi(tender.City),
		Organization:    ConvertOrganizationModelToApi(tender.Organization),
		Price:           float64(tender.Price / 100),
		IsContractPrice: tender.IsContractPrice,
		IsNdsPrice:      tender.IsNDSPrice,
		IsDraft:         tender.IsDraft,
		FloorSpace:      tender.FloorSpace,
		Description:     tender.Description,
		Wishes:          tender.Wishes,
		Specification:   api.URL(tender.Specification),
		Attachments: convert.Slice[[]string, []api.URL](
			tender.Attachments, func(u string) api.URL { return api.URL(u) },
		),
		Services: convert.Slice[[]Service, api.Services](
			tender.Services, ConvertServiceModelToApi,
		),
		Objects: convert.Slice[[]Object, api.Objects](
			tender.Objects, ConvertObjectModelToApi,
		),
		Status:             "",
		VerificationStatus: api.OptVerificationStatus{Value: tender.VerificationStatus.ToAPI(), Set: tender.VerificationStatus != 0},
		ReceptionStart:     tender.ReceptionStart,
		ReceptionEnd:       tender.ReceptionEnd,
		WorkStart:          tender.WorkStart,
		WorkEnd:            tender.WorkEnd,
		CreatedAt:          tender.CreatedAt,
		UpdatedAt:          tender.UpdatedAt,
	}

	if tender.WinnerOrganization.Set {
		tenderApi.WinnerOrganization = api.OptOrganization{
			Value: ConvertOrganizationModelToApi(tender.WinnerOrganization.Value),
			Set:   true}
	}

	return tenderApi
}

func ConvertServiceModelToApi(service Service) api.Service {
	return api.Service{
		ID:       service.ID,
		ParentID: api.OptInt{Value: service.ParentID, Set: service.ParentID != 0},
		Name:     service.Name,
	}
}

func ConvertObjectModelToApi(object Object) api.Object {
	return api.Object{
		ID:       object.ID,
		ParentID: api.OptInt{Value: object.ParentID, Set: object.ParentID != 0},
		Name:     object.Name,
	}
}
