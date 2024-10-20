package store

import "time"

type TenderCreateParams struct {
	Name            string
	Price           float64
	IsContractPrice bool
	IsNDSPrice      bool
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

type TenderServicesCreateParams struct {
	TenderID    int
	ServicesIDs []int
}

type TenderObjectsCreateParams struct {
	TenderID   int
	ObjectsIDs []int
}
