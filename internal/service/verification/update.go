package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error {
	err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		result, err := s.verificationStore.UpdateStatus(ctx, qe, store.VerificationRequestUpdateStatusParams{
			UserID:        params.UserID,
			RequestID:     params.RequesID,
			Status:        params.Status,
			ReviewComment: params.ReviewComment,
		})
		if err != nil {
			return fmt.Errorf("update request status: %w", err)
		}

		var status models.TenderStatus

		if params.Status == models.VerificationStatusApproved {
			status = models.ReceptionStatus
		} else if params.Status == models.VerificationStatusDeclined {
			status = models.RemovedByModeratorStatus
		} else {
			status = models.InvalidStatus
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
				Status:             status,
			})

		case models.ObjectTypeAddition:
			err = s.additionStore.UpdateVerificationStatus(ctx, qe, store.AdditionUpdateVerifStatusParams{
				AdditionID:         result.ObjectID,
				VerificationStatus: params.Status,
			})

		case models.ObjectTypeQuestionAnswer:
			err = s.questionAnswerStore.UpdateVerificationStatus(ctx, qe, store.QuestionAnswerVerifStatusUpdateParams{
				QuestionAnswerID:   result.ObjectID,
				VerificationStatus: params.Status,
			})
		default:
			return fmt.Errorf("invalid object type: %v", result.ObjectType)
		}
		if err != nil {
			return fmt.Errorf("update object type=%v status: %w", result.ObjectType, err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("run transaction: %w", err)
	}
	return nil
}
