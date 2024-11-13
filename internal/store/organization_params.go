package store

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
