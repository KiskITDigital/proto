package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error {
	return s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		Brand: params.Brand,
		AvatarURL: params.AvatarURL,
	})
}
