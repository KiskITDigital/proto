package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (int, error) {
	builder := squirrel.Update("tenders")

	if params.Name.Set {
		builder = builder.Set("name", params.Name.Value)
	}
	if params.ServiceIDs.Set {
		builder = builder.Set("services_ids", pq.Array(params.ServiceIDs.Value))
	}
	if params.ObjectIDs.Set {
		builder = builder.Set("objects_ids", pq.Array(params.ObjectIDs.Value))
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
			RETURNING id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var id int

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}
