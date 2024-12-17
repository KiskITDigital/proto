package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

// Избранное
func (h *Handler) V1OrganizationsFavouritesFavouriteIDDelete(ctx context.Context, params api.V1OrganizationsFavouritesFavouriteIDDeleteParams) (api.V1OrganizationsFavouritesFavouriteIDDeleteRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDFavouritesGet(ctx context.Context, params api.V1OrganizationsOrganizationIDFavouritesGetParams) (api.V1OrganizationsOrganizationIDFavouritesGetRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}

func (h *Handler) V1OrganizationsOrganizationIDFavouritesPost(ctx context.Context, req *api.V1OrganizationsOrganizationIDFavouritesPostReq, params api.V1OrganizationsOrganizationIDFavouritesPostParams) (api.V1OrganizationsOrganizationIDFavouritesPostRes, error) {
	return nil, cerr.Wrap(fmt.Errorf("not impl"), cerr.CodeInternal, "func not impl", nil)
}
