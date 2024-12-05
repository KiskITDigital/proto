package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type OrganizationGetParams struct {
	IsContractor models.Optional[bool]
	Offset       uint64
	Limit        uint64
}

type OrganizationUpdateBrandParams struct {
	OrganizationID int
	Brand          models.Optional[string]
	AvatarURL      models.Optional[string]
}
