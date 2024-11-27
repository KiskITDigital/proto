package verification

import (
	"context"
	"fmt"

	// "gitlab.ubrato.ru/ubrato/core/internal/service"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error {
	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		result, err := s.verificationStore.UpdateStatus(ctx, qe, store.VerificationRequestUpdateStatusParams{
			UserID:    params.UserID,
			RequestID: params.RequesID,
			Status:    params.Status,
		})
		if err != nil {
			return fmt.Errorf("update request status: %w", err)
		}

		switch result.ObjectType {
		case models.ObjectTypeOrganization:
			err = s.organizationStore.UpdateVerificationStatus(ctx, qe, store.OrganizationUpdateVerifStatusParams{
				OrganizationID:     result.ObjectID,
				VerificationStatus: params.Status,
			})

		case models.ObjectTypeTender:
			err = s.tenderStore.UpdateVerificationStatus(ctx, qe, store.TenderUpdateVerifStatusParams{
				TenderID:           result.ObjectID,
				VerificationStatus: params.Status,
			})

		case models.ObjectTypeComment:
			// TODO: add update comment status
			return fmt.Errorf("update comment type not impl")

		default:
			return fmt.Errorf("invalid object type: %v", result.ObjectType)
		}
		if err != nil {
			return fmt.Errorf("update object type=%v status: %w", result.ObjectType, err)
		}

		return nil
	})
}
