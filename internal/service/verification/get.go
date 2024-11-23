package verification

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	storeParams := store.VerificationRequestsObjectGetParams{
		ObjectType: params.ObjectType,
		Status:     params.Status,
		Offset:     params.Offset,
		Limit:      params.Limit,
	}

	switch params.ObjectType {
	case models.ObjectTypeOrganization:
		return s.verificationStore.GetOrganizationRequests(ctx, s.psql.DB(), storeParams)

	case models.ObjectTypeTender:
		// 1. get request with tenderID
		requests, err := s.verificationStore.GetTendersRequests(ctx, s.psql.DB(), storeParams)
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
		return nil, errors.New("get comments request not impl")
	}

	return nil, fmt.Errorf("invalid object type: %v", params.ObjectType)
}
