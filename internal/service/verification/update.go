package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error {
	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		result, err := s.verificationStore.UpdateStatus(ctx, qe, store.VerificationRequestUpdateStatusParams{
			UserID:        params.UserID,
			RequestID:     params.RequesID,
			Status:        params.Status,
			ReviewComment: params.ReviewComment,
		})
		if err != nil {
			return fmt.Errorf("update request status: %w", err)
		}

		notification := &modelsv1.Notification{
			User: &modelsv1.NotifiedUser{},
			Verification: &modelsv1.Verification{
				Status:  modelsv1.Status(params.Status),
				Comment: params.ReviewComment.Value,
			},
			Object: &modelsv1.Object{
				Id:     int32(result.ObjectID),
				Type:   modelsv1.ObjectType(result.ObjectType),
				Tender: &modelsv1.Tender{},
			},
		}

		// Статусы
		var status models.TenderStatus
		if params.Status == models.VerificationStatusApproved {
			status = models.ReceptionStatus
		} else if params.Status == models.VerificationStatusDeclined {
			status = models.RemovedByModeratorStatus
		} else {
			status = models.InvalidStatus
		}

		var (
			topic              broker.Topic
			userOrganizationID int
		)

		switch result.ObjectType {
		case models.ObjectTypeOrganization:
			topic = broker.UbratoOrganizationVerification

			err = s.organizationStore.UpdateVerificationStatus(ctx, qe, store.OrganizationUpdateVerifStatusParams{
				OrganizationID:     result.ObjectID,
				VerificationStatus: params.Status,
			})

			isContractor, err := s.organizationStore.GetIsContractorByID(ctx, qe, result.ObjectID)
			if err != nil {
				return fmt.Errorf("get organization is_contractor by id: %w", err)
			}
			notification.User.IsContractor = isContractor
			userOrganizationID = result.ObjectID

		case models.ObjectTypeTender:
			topic = broker.UbratoTenderVerification

			err = s.tenderStore.UpdateVerificationStatus(ctx, qe, store.TenderUpdateVerifStatusParams{
				TenderID:           result.ObjectID,
				VerificationStatus: params.Status,
				Status:             status,
			})

			tenderNotifyInfo, err := s.tenderStore.GetTenderNotifyInfoByID(ctx, qe, result.ObjectID)
			if err != nil {
				return fmt.Errorf("get tender name, receptionStart by id=%v: %w", result.ObjectID, err)
			}
			notification.Object.Tender.Title = tenderNotifyInfo.Name
			notification.Object.Tender.ReceptionStart = timestamppb.New(tenderNotifyInfo.ReceptionStart)
			userOrganizationID = tenderNotifyInfo.Organization.ID

		case models.ObjectTypeAddition:
			topic = broker.UbratoTenderAdditionVerification

			err = s.additionStore.UpdateVerificationStatus(ctx, qe, store.AdditionUpdateVerifStatusParams{
				AdditionID:         result.ObjectID,
				VerificationStatus: params.Status,
			})

		case models.ObjectTypeQuestionAnswer:
			topic = broker.UbratoTenderQuestionAnswerVerification

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

		userID, err := s.userStore.GetUserIDByOrganizationID(ctx, qe, userOrganizationID)
		if err != nil {
			return fmt.Errorf("get userID by orgID: %w", err)
		}
		notification.User.Id = int32(userID)

		// Уведомления
		b, err := proto.Marshal(&eventsv1.SentNotification{Notification: notification})
		if err != nil {
			return fmt.Errorf("marhal notification proto: %w", err)
		}

		err = s.broker.Publish(ctx, topic, b)
		if err != nil {
			return fmt.Errorf("notification: %w", err)
		}

		return nil
	})
}
