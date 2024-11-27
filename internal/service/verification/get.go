package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"

	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, requestID int) (models.VerificationRequest[models.VerificationObject], error) {
	request, err := s.verificationStore.GetByIDWithEmptyObject(ctx, s.psql.DB(), requestID)
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get req by id=%v: %w", requestID, err)
	}

	var object models.VerificationObject
	switch request.ObjectType {
	case models.ObjectTypeOrganization:
		object, err = s.organizationStore.GetByID(ctx, s.psql.DB(), request.ObjectID)

	case models.ObjectTypeTender:
		object, err = s.tenderStore.GetByID(ctx, s.psql.DB(), request.ObjectID)

	case models.ObjectTypeComment:
		object, err = s.commentStore.GetByID(ctx, s.psql.DB(), request.ObjectID)
	}
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get object type=%v by id=%v: %w", request.ObjectType, request.ObjectID, err)
	}

	request.Object = object

	return request, nil
}

func (s *Service) Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	storeParams := store.VerificationRequestsObjectGetParams{
		ObjectType: models.NewOptional(params.ObjectType),
		Status:     params.Status,
		Offset:     params.Offset,
		Limit:      params.Limit,
	}

	switch params.ObjectType {
	case models.ObjectTypeOrganization:
		return s.verificationStore.GetOrganizationRequests(ctx, s.psql.DB(), storeParams)

	case models.ObjectTypeTender:
		// 1. get request with tenderID
		requests, err := s.verificationStore.GetWithEmptyObject(ctx, s.psql.DB(), storeParams)
		if err != nil {
			return nil, fmt.Errorf("get tenders req: %w", err)
		}

		var tenderIDs []int
		for _, req := range requests {
			tenderIDs = append(tenderIDs, req.ObjectID)
		}

		// 2. get tenders by tenderIDs
		tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
			TenderIDs: models.Optional[[]int]{Value: tenderIDs, Set: true}})
		if err != nil {
			return nil, fmt.Errorf("get tenders: %w", err)
		}

		tenderMap := make(map[int]models.Tender)
		for _, tender := range tenders {
			tenderMap[tender.ID] = tender
		}

		for i := range requests {
			if tender, ok := tenderMap[requests[i].ObjectID]; ok {
				requests[i].Object = tender
			}
		}

		return requests, nil

	case models.ObjectTypeComment:
		return s.verificationStore.GetCommentRequests(ctx, s.psql.DB(), storeParams)
	}

	return nil, fmt.Errorf("invalid object type: %v", params.ObjectType)
}
