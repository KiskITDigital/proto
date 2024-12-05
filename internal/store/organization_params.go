package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type OrganizationGetParams struct {
	IsContractor models.Optional[bool]
	Offset       uint64
	Limit        uint64
}

type OrganizationCreateParams struct {
	BrandName    string
	FullName     string
	ShortName    string
	IsContractor bool
	INN          string
	OKPO         string
	OGRN         string
	KPP          string
	TaxCode      string
	Address      string
}

type OrganizationUpdateParams struct {
	OrganizationID int
	Brand          models.Optional[string]
	AvatarURL      models.Optional[string]
}

type OrganizationAddUserParams struct {
	OrganizationID int
	UserID         int
	IsOwner        bool
}

type OrganizationUpdateVerifStatusParams struct {
	OrganizationID     int
	VerificationStatus models.VerificationStatus
}
