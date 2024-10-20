package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	tenderService "gitlab.ubrato.ru/ubrato/core/internal/service/tender"
)

func (h *Handler) V1TendersCreatePost(ctx context.Context, req *api.V1TendersCreatePostReq) (api.V1TendersCreatePostRes, error) {
	tender, err := h.svc.Create(ctx, tenderService.CreateParams{
		Name:            req.GetName(),
		CityID:          req.GetCity(),
		Price:           int(req.GetPrice() * 100),
		IsContractPrice: req.GetIsContractPrice(),
		IsNDSPrice:      req.GetIsNdsPrice(),
		IsDraft:         req.GetIsDraft().Value,
		FloorSpace:      req.GetFloorSpace(),
		Description:     req.GetDescription().Value,
		Wishes:          req.GetWishes().Value,
		Specification:   string(req.Specification.Value),
		Attachments: convert.Slice[[]api.URL, []string](
			req.GetAttachments(), func(u api.URL) string { return string(u) },
		),
		ServiceIDs:     req.GetServices(),
		ObjectIDs:      req.GetObjects(),
		ReceptionStart: req.GetReceptionStart(),
		ReceptionEnd:   req.GetReceptionEnd(),
		WorkStart:      req.GetWorkStart(),
		WorkEnd:        req.GetWorkEnd(),
	})
	if err != nil {
		return nil, fmt.Errorf("create tender: %w", err)
	}

	return &api.V1TendersCreatePostCreated{
		Data: api.V1TendersCreatePostCreatedData{
			Tender: models.ConvertTenderModelToApi(tender),
		},
	}, nil
}
