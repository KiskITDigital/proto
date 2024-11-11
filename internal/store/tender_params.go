package store

import (
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type TenderCreateParams struct {
	Name            string
	Price           int
	IsContractPrice bool
	IsNDSPrice      bool
	IsDraft         bool
	CityID          int
	FloorSpace      int
	Description     string
	Wishes          string
	Specification   string
	Attachments     []string
	ReceptionStart  time.Time
	ReceptionEnd    time.Time
	WorkStart       time.Time
	WorkEnd         time.Time
	OrganizationID  int
}

type TenderGetParams struct {
	OrganizationID models.Optional[int]
	WithDrafts     bool
}

type TenderUpdateParams struct {
	ID              int
	Name            models.Optional[string]
	Price           models.Optional[int]
	IsContractPrice models.Optional[bool]
	IsNDSPrice      models.Optional[bool]
	IsDraft         models.Optional[bool]
	CityID          models.Optional[int]
	FloorSpace      models.Optional[int]
	Description     models.Optional[string]
	Wishes          models.Optional[string]
	Specification   models.Optional[string]
	Attachments     models.Optional[[]string]
	ReceptionStart  models.Optional[time.Time]
	ReceptionEnd    models.Optional[time.Time]
	WorkStart       models.Optional[time.Time]
	WorkEnd         models.Optional[time.Time]
}

type TenderServicesCreateParams struct {
	TenderID    int
	ServicesIDs []int
}

type TenderObjectsCreateParams struct {
	TenderID   int
	ObjectsIDs []int
}

type TenderObjectsDeleteParams struct {
	TenderID   int
	ObjectsIDs []int
}

type TenderServicesDeleteParams struct {
	TenderID    int
	ServicesIDs []int
}

type TenderCreateResponseParams struct {
	TenderID       int
	OrganizationID int
	Price          int
	IsNds          bool
}
