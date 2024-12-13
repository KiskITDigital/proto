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
	requests, err := s.verificationStore.GetWithEmptyObject(ctx, s.psql.DB(), store.VerificationRequestsObjectGetParams{
		ObjectType: models.NewOptional(params.ObjectType),
		ObjectID:   params.ObjectID,
		Status:     params.Status,
		Offset:     params.Offset,
		Limit:      params.Limit,
	})
	if err != nil {
		return nil, fmt.Errorf("get object req: %w", err)
	}

	var objectIDs []int
	for _, req := range requests {
		objectIDs = append(objectIDs, req.ObjectID)
	}

	switch params.ObjectType {
	case models.ObjectTypeOrganization:
		organizations, err := s.organizationStore.Get(ctx, s.psql.DB(), store.OrganizationGetParams{
			OrganizationIDs: objectIDs})
		if err != nil {
			return nil, fmt.Errorf("get tenders: %w", err)
		}

		organizationMap := make(map[int]models.Organization)
		for _, organization := range organizations {
			organizationMap[organization.ID] = organization
		}

		for i := range requests {
			if organization, ok := organizationMap[requests[i].ObjectID]; ok {
				requests[i].Object = organization
			}
		}

	case models.ObjectTypeTender:
		tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
			TenderIDs: models.Optional[[]int]{Value: objectIDs, Set: true}})
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

	case models.ObjectTypeComment:
		// return s.verificationStore.GetCommentRequests(ctx, s.psql.DB(), storeParams)
	default:
		return nil, fmt.Errorf("invalid object type: %v", params.ObjectType)
	}

	return requests, nil
}
