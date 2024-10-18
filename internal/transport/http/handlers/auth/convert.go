package auth

import (
	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func ConvertUserModelToApi(user models.User) api.User {
	return api.User{
		ID:            user.ID,
		Email:         api.Email(user.Email),
		Phone:         api.Phone(user.Phone),
		FirstName:     api.Name(user.FirstName),
		LastName:      api.Name(user.LastName),
		MiddleName:    api.Name(user.MiddleName),
		AvatarURL:     api.URL(user.AvatarURL),
		Verified:      user.Verified,
		EmailVerified: user.EmailVerified,
		Role:          api.Role(user.Role),
		IsContractor:  user.IsContractor,
		IsBanned:      user.IsBanned,
		Organization: api.OptOrganization{
			Value: api.Organization(ConvertOrganizationModelToApi(user.Organization)),
			Set:   user.Organization.ID != 0,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ConvertOrganizationModelToApi(organization models.Organization) api.Organization {
	return api.Organization{
		ID:        organization.ID,
		BrandName: api.Name(organization.BrandName),
		FullName:  api.Name(organization.FullName),
		ShortName: api.Name(organization.ShortName),
		Inn:       api.Inn(organization.INN),
		Okpo:      api.Okpo(organization.OKPO),
		Ogrn:      api.Ogrn(organization.OGRN),
		Kpp:       api.Kpp(organization.KPP),
		TaxCode:   api.TaxCode(organization.TaxCode),
		Address:   organization.Address,
		AvatarURL: api.URL(organization.AvatarURL),
		Emails: convert.Slice[models.ContactInfos, []api.ContactInfo](
			organization.Emails, ConvertContactInfoModelToApi,
		),
		Phones: convert.Slice[models.ContactInfos, []api.ContactInfo](
			organization.Phones, ConvertContactInfoModelToApi,
		),
		Messengers: convert.Slice[models.ContactInfos, []api.ContactInfo](
			organization.Messengers, ConvertContactInfoModelToApi,
		),
		CreatedAt: organization.CreatedAt,
		UpdatedAt: organization.UpdatedAt,
	}
}

func ConvertContactInfoModelToApi(info models.ContactInfo) api.ContactInfo {
	return api.ContactInfo{
		Contact: info.Info,
		Info:    info.Info,
	}
}
