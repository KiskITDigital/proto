package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type Portfolio struct {
	ID             int
	OrganizationID int
	Title          string
	Description    string
	Attachments    []string
	CreatedAt      time.Time
	UpdatedAt      Optional[time.Time]
}

func ConvertPortfolioModelToApi(portfolio Portfolio) api.Portfolio {
	return api.Portfolio{
		ID:          portfolio.ID,
		Title:       portfolio.Title,
		Description: api.Description(portfolio.Description),
		Attachments: convert.Slice[[]string, []api.URL](
			portfolio.Attachments, func(u string) api.URL { return api.URL(u) },
		),
		CreatedAt: portfolio.CreatedAt,
		UpdatedAt: api.OptDateTime{Value: portfolio.UpdatedAt.Value, Set: portfolio.UpdatedAt.Set},
	}
}
