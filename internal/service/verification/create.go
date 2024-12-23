package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.VerificationRequestCreateParams) error {
	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		err := s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:    params.ObjectID,
			ObjectType:  params.ObjectType,
			Attachments: params.Attachments,
		})
		if err != nil {
			return fmt.Errorf("create verification: %w", err)
		}

		err = s.organizationStore.UpdateVerificationStatus(ctx, qe, store.OrganizationUpdateVerifStatusParams{
			OrganizationID:     params.ObjectID,
			VerificationStatus: models.VerificationStatusInReview,
		})
		if err != nil {
			return fmt.Errorf("update organization verification status: %w", err)
		}

		return err
	})
}
