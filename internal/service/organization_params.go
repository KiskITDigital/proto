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

type OrganizationUpdateContactsParams struct {
	OrganizationID int
	Emails         models.Optional[models.ContactInfos]
	Phones         models.Optional[models.ContactInfos]
	Messengers     models.Optional[models.ContactInfos]
}

type OrganizationCreateVerificationRequestParams struct {
	OrganizationID int
	Attachments    models.Attachments
}

type OrganizationUpdateContractorParams struct {
	OrganizationID int
	Description    models.Optional[string]
	Cities         models.Optional[[]int]
	Services       models.Optional[[]int]
	Objects        models.Optional[[]int]
}

type OrganizationUpdateCustomerParams struct {
	OrganizationID int
	Description    models.Optional[string]
	CityIDs        []int
}
