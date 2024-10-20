package tender

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (models.Tender, error) {
	builder := squirrel.Update("tenders")

	if params.Name.Set {
		builder = builder.Set("name", params.Name.Value)
	}
	if params.Price.Set {
		builder = builder.Set("price", params.Price.Value)
	}
	if params.IsContractPrice.Set {
		builder = builder.Set("is_contract_price", params.IsContractPrice.Value)
	}
	if params.IsNDSPrice.Set {
		builder = builder.Set("is_nds_price", params.IsNDSPrice.Value)
	}
	if params.IsDraft.Set {
		builder = builder.Set("is_draft", params.IsDraft.Value)
	}
	if params.CityID.Set {
		builder = builder.Set("city_id", params.CityID.Value)
	}
	if params.FloorSpace.Set {
		builder = builder.Set("floor_space", params.FloorSpace.Value)
	}
	if params.Description.Set {
		builder = builder.Set("description", params.Description.Value)
	}
	if params.Wishes.Set {
		builder = builder.Set("wishes", params.Wishes.Value)
	}
	if params.Specification.Set {
		builder = builder.Set("specification", params.Specification.Value)
	}
	if params.Attachments.Set {
		builder = builder.Set("attachments", pq.Array(params.Attachments.Value))
	}
	if params.ReceptionStart.Set {
		builder = builder.Set("reception_start", params.ReceptionStart.Value)
	}
	if params.ReceptionEnd.Set {
		builder = builder.Set("reception_end", params.ReceptionEnd.Value)
	}
	if params.WorkStart.Set {
		builder = builder.Set("work_start", params.WorkStart.Value)
	}
	if params.WorkEnd.Set {
		builder = builder.Set("work_end", params.WorkEnd.Value)
	}

	builder = builder.
		Where(squirrel.Eq{"id": params.ID}).
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

	var (
		tender      models.Tender
		description sql.NullString
		wishes      sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&tender.ID,
		&tender.Name,
		&tender.Price,
		&tender.IsContractPrice,
		&tender.IsNDSPrice,
		&tender.IsDraft,
		&tender.City.ID,
		&tender.FloorSpace,
		&description,
		&wishes,
		&tender.Specification,
		pq.Array(&tender.Attachments),
		&tender.Verified,
		&tender.ReceptionStart,
		&tender.ReceptionEnd,
		&tender.WorkStart,
		&tender.WorkEnd,
		&tender.Organization.ID,
	)
	if err != nil {
		return models.Tender{}, fmt.Errorf("query row: %w", err)
	}

	tender.Description = description.String
	tender.Wishes = wishes.String

	return tender, nil
}
