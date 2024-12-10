package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

const (
	Page    = 0
	PerPage = 100
)

type Pagination struct {
	Page    uint64
	Pages   uint64
	PerPage uint64
}

func ConvertPaginationToAPI(p Pagination) api.Pagination {
	return api.Pagination{
		Page:    int(p.Page),
		Pages:   int(p.Pages),
		PerPage: int(p.PerPage),
	}
}
