package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (models.Tender, error) {
	builder := squirrel.
		Insert("tenders").
		Columns(
			"name",
			"price",
			"is_contract_price",
			"is_nds_price",
			"is_draft",
			"city_id",
			"floor_space",
			"description",
			"wishes",
			"specification",
			"attachments",
			"reception_start",
			"reception_end",
			"work_start",
			"work_end",
			"organization_id",
		).
		Values(
			params.Name,
			params.Price,
			params.IsContractPrice,
			params.IsNDSPrice,
			params.IsDraft,
			params.CityID,
			params.FloorSpace,
			params.Description,
			params.Wishes,
			params.Specification,
			pq.Array(params.Attachments),
			params.ReceptionStart,
			params.ReceptionEnd,
			params.WorkStart,
			params.WorkEnd,
			params.OrganizationID,
		).
		Suffix(`
			RETURNING
				id,
				name,
				price,
				is_contract_price,
				is_nds_price,
				is_draft,
				city_id,
				floor_space,
				description,
				wishes,
				specification,
				attachments,
				verified,
				reception_start,
				reception_end,
				work_start,
				work_end,
				organization_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdTender models.Tender

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdTender.ID,
		&createdTender.Name,
		&createdTender.Price,
		&createdTender.IsContractPrice,
		&createdTender.IsNDSPrice,
		&createdTender.IsDraft,
		&createdTender.City.ID,
		&createdTender.FloorSpace,
		&createdTender.Description,
		&createdTender.Wishes,
		&createdTender.Specification,
		pq.Array(&createdTender.Attachments),
		&createdTender.Verified,
		&createdTender.ReceptionStart,
		&createdTender.ReceptionEnd,
		&createdTender.WorkStart,
		&createdTender.WorkEnd,
		&createdTender.Organization.ID,
	)
	if err != nil {
		return models.Tender{}, fmt.Errorf("query row: %w", err)
	}

	return createdTender, nil
}
