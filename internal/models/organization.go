package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type ContactInfo struct {
	Contact string `json:"contact"`
	Info    string `json:"info"`
}

type ContactInfos []ContactInfo

func (a ContactInfos) Value() (driver.Value, error) {
	if a == nil {
		return []byte("[]"), nil
	}

	return json.Marshal(a)
}

func (a *ContactInfos) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Organization struct {
	VerificationObject

	ID                 int
	BrandName          string
	FullName           string
	ShortName          string
	INN                string
	OKPO               string
	OGRN               string
	KPP                string
	TaxCode            string
	Address            string
	AvatarURL          string
	VerificationStatus VerificationStatus
	IsContractor       bool
	IsBanned           bool
	Emails             ContactInfos
	Phones             ContactInfos
	Messengers         ContactInfos
	CustomerInfo       CustomerInfo
	ContractorInfo     ContractorInfo
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func (o Organization) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:         api.OrganizationVerificationRequestObject,
		Organization: ConvertOrganizationModelToApi(o),
	}
}

func ConvertOrganizationModelToApi(organization Organization) api.Organization {
	return api.Organization{
		ID:                 organization.ID,
		BrandName:          api.Name(organization.BrandName),
		FullName:           api.Name(organization.FullName),
		ShortName:          api.Name(organization.ShortName),
		VerificationStatus: api.NewOptVerificationStatus(organization.VerificationStatus.ToAPI()),
		IsContractor:       organization.IsContractor,
		IsBanned:           organization.IsBanned,
		Inn:                api.Inn(organization.INN),
		Okpo:               api.Okpo(organization.OKPO),
		Ogrn:               api.Ogrn(organization.OGRN),
		Kpp:                api.Kpp(organization.KPP),
		TaxCode:            api.TaxCode(organization.TaxCode),
		Address:            organization.Address,
		AvatarURL:          api.OptURL{Value: api.URL(organization.AvatarURL), Set: organization.AvatarURL != ""},
		Emails: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Emails, ConvertContactInfoModelToApi,
		),
		Phones: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Phones, ConvertContactInfoModelToApi,
		),
		Messengers: convert.Slice[ContactInfos, []api.ContactInfo](
			organization.Messengers, ConvertContactInfoModelToApi,
		),
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}
}

func ConvertContactInfoModelToApi(info ContactInfo) api.ContactInfo {
	return api.ContactInfo{
		Contact: info.Info,
		Info:    info.Info,
	}
}

func ConvertAPIToContactInfo(info api.ContactInfo) ContactInfo {
	return ContactInfo{
		Contact: info.Contact,
		Info:    info.Info,
	}
}

type CustomerInfo struct{}

func (a CustomerInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *CustomerInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type ContractorInfo struct{}

func (a ContractorInfo) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *ContractorInfo) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
