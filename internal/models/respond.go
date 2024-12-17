package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type RespondPagination struct {
	Responds   []Respond
	Pagination pagination.Pagination
}

type Respond struct {
	TenderID       int
	OrganizationID int
	Price          int
	IsNDSPrice     bool
	CreatedAt      time.Time
}

func ConvertRespondModelToApi(respond Respond) api.Respond {
	return api.Respond{
		TenderID:       respond.TenderID,
		OrganizationID: respond.OrganizationID,
		Price:          respond.Price,
		IsNdsPrice:     respond.IsNDSPrice,
		CreatedAt:      respond.CreatedAt,
	}
}
