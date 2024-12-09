package organization

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1OrganizationsOrganizationIDProfileBrandPut(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDProfileBrandPutReq,
	params api.V1OrganizationsOrganizationIDProfileBrandPutParams,
) (api.V1OrganizationsOrganizationIDProfileBrandPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	if err := h.organizationService.UpdateBrand(ctx, service.OrganizationUpdateBrandParams{
		OrganizationID: params.OrganizationID,
		Brand:          models.Optional[string]{Value: req.GetBrand().Value, Set: req.GetBrand().Set},
		AvatarURL:      models.Optional[string]{Value: string(req.GetAvatarURL().Value), Set: req.GetAvatarURL().Set},
	}); err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileBrandPutOK{}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileContactsPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContactsPutReq, params api.V1OrganizationsOrganizationIDProfileContactsPutParams) (api.V1OrganizationsOrganizationIDProfileContactsPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	if err := h.organizationService.UpdateContacts(ctx, service.OrganizationUpdateContactsParams{
		OrganizationID: params.OrganizationID,
		Emails: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetEmails(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetEmails()) != 0,
		},
		Phones: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetPhones(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetPhones()) != 0,
		},
		Messengers: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetMessengers(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetMessengers()) != 0,
		},
	}); err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization contacts: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileContactsPutOK{}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileCustomerPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileCustomerPutReq, params api.V1OrganizationsOrganizationIDProfileCustomerPutParams) (api.V1OrganizationsOrganizationIDProfileCustomerPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	organization, err := h.organizationService.UpdateCustomer(ctx, service.OrganizationUpdateCustomerParams{
		OrganizationID: params.OrganizationID,
		Description:    models.Optional[string]{Value: string(req.GetDescription().Value), Set: req.GetDescription().Set},
		CityIDs:        req.GetCityIds()})
	if err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileCustomerPutOK{
		Data: models.ConvertOrganizationModelToApi(organization),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileContractorPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContractorPutReq, params api.V1OrganizationsOrganizationIDProfileContractorPutParams) (api.V1OrganizationsOrganizationIDProfileContractorPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	organization, err := h.organizationService.UpdateContractor(ctx, service.OrganizationUpdateContractorParams{
		OrganizationID: params.OrganizationID,
		Description:    models.Optional[string]{Value: string(req.GetDescription().Value), Set: req.GetDescription().Set},
		CityIDs:        req.GetCityIds(),
		ServiceIDs:     req.GetServiceIds(),
		ObjectIDs:      req.GetObjectsIds(),
	})
	if err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileContractorPutOK{
		Data: models.ConvertOrganizationModelToApi(organization),
	}, nil
}
