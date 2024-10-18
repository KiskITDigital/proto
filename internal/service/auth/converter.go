package auth

import (
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func convertFindByINNResponseToOrganization(resp dadata.FindByInnResponse) (models.Organization, error) {
	if len(resp.Suggestions) == 0 {
		return models.Organization{}, fmt.Errorf("no suggestions")
	}

	suggestion := resp.Suggestions[0]

	return models.Organization{
		BrandName: suggestion.Data.Name.Short,
		FullName:  suggestion.Data.Name.FullWithOpf,
		ShortName: suggestion.Data.Name.ShortWithOpf,
		INN:       suggestion.Data.INN,
		OKPO:      suggestion.Data.OKPO,
		ORGN:      suggestion.Data.ORGN,
		KPP:       suggestion.Data.KPP,
		TaxCode:   suggestion.Data.Address.Data.TaxOffice,
		Address:   suggestion.Data.Address.UnrestrictedValue,
	}, nil
}
