package favourite

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.FavouriteCreateParams) error {
	switch params.ObjectType {
	case models.FavouriteTypeOrganization:
		organization, err := s.organizationStore.GetByID(ctx, s.psql.DB(), params.ObjectID)
		if err != nil {
			return fmt.Errorf("get organization by id: %w", err)
		}

		if organization.VerificationStatus != models.VerificationStatusApproved || !organization.IsContractor {
			return errors.New("cannot add organization")
		}

	case models.FavouriteTypeTender:
		tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.ObjectID)
		if err != nil {
			return fmt.Errorf("get tender by id: %w", err)
		}

		if tender.VerificationStatus != models.VerificationStatusApproved || tender.IsDraft {
			return errors.New("cannot add tender")
		}

	default:
		return errors.New("unsupported object type")
	}

	return s.favouriteStore.Create(ctx, s.psql.DB(), store.FavouriteCreateParams{
		OrganizationID: params.OrganizationID,
		ObjectType:     params.ObjectType,
		ObjectID:       params.ObjectID,
	})
}
