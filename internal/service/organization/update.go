package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error {
	return s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		Brand:          params.Brand,
		AvatarURL:      params.AvatarURL,
	})
}

func (s *Service) UpdateContacts(ctx context.Context, params service.OrganizationUpdateContactsParams) error {
	return s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		Emails:         params.Emails,
		Phones:         params.Phones,
		Messengers:     params.Messengers,
	})
}
