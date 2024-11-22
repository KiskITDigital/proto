package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type OrganizationGetParams struct{}

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

type OrganizationAddUserParams struct {
	OrganizationID int
	UserID         int
	IsOwner        bool
}

type OrganizationUpdateVerifStatusParams struct {
	OrganizationID     int
	VerificationStatus models.VerificationStatus
}
