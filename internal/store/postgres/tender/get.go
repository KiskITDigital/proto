package tender

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/deduplicate"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error) {
	tenders, err := s.List(ctx, qe, store.TenderListParams{
		TenderIDs: models.Optional[[]int]{Set: true, Value: []int{id}},
	})
	if err != nil {
		return models.Tender{}, fmt.Errorf("list tenders: %w", err)
	}

	if len(tenders) == 0 {
		return models.Tender{}, errors.New("tender not found")
	}

	return tenders[0], nil
}

func (s *TenderStore) List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error) {
	builder := squirrel.
		Select(
			"t.id",
			"t.city_id",
			"t.services_ids",
			"t.objects_ids",
			"t.name",
			"t.price",
			"t.is_contract_price",
			"t.is_nds_price",
			"t.floor_space",
			"t.description",
			"t.wishes",
			"t.specification",
			"t.attachments",
			"t.status",
			"t.verification_status",
			"t.is_draft",
			"t.reception_start",
			"t.reception_end",
			"t.work_start",
			"t.work_end",
			"t.created_at",
			"t.updated_at",
			"c.name",
			"c.id",
			"r.name",
			"r.id",
			"o.id",
			"o.brand_name",
			"o.full_name",
			"o.short_name",
			"o.inn",
			"o.okpo",
			"o.ogrn",
			"o.kpp",
			"o.tax_code",
			"o.address",
			"o.avatar_url",
			"o.emails",
			"o.phones",
			"o.messengers",
			"o.verification_status",
			"o.is_contractor",
			"o.is_banned",
			"o.customer_info",
			"o.contractor_info",
			"o.created_at",
			"o.updated_at",
			"ow.id",
			"ow.brand_name",
			"ow.full_name",
			"ow.short_name",
			"ow.inn",
			"ow.okpo",
			"ow.ogrn",
			"ow.kpp",
			"ow.tax_code",
			"ow.address",
			"ow.avatar_url",
			"ow.emails",
			"ow.phones",
			"ow.messengers",
			"ow.verification_status",
			"ow.is_contractor",
			"ow.is_banned",
			"ow.customer_info",
			"ow.contractor_info",
			"ow.created_at",
			"ow.updated_at",
		).
		From("tenders AS t").
		Join("cities AS c ON c.id = t.city_id").
		Join("regions AS r ON r.id = c.region_id").
		Join("organizations AS o ON o.id = t.organization_id").
		LeftJoin("organizations AS ow ON ow.id = t.winner_organization_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.OrganizationID.Set {
		builder = builder.Where(squirrel.Eq{"t.organization_id": params.OrganizationID.Value})
	}

	if params.TenderIDs.Set {
		builder = builder.Where(squirrel.Eq{"t.id": params.TenderIDs.Value})
	}

	if !params.WithDrafts {
		builder = builder.Where(squirrel.Eq{"t.is_draft": false})
	}

	if params.VerifiedOnly {
		builder = builder.Where(squirrel.Eq{"t.verification_status": models.VerificationStatusApproved})
	}

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	var tenders []models.Tender
	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var (
		tenderToServices = make(map[int][]int)
		tenderToObjects  = make(map[int][]int)
		serviceIDs       = make([]int, 0)
		objectIDs        = make([]int, 0)
	)

	for rows.Next() {
		var (
			tender           models.Tender
			tenderServiceIDs pq.Int64Array
			tenderObjectIDs  pq.Int64Array
			description      sql.NullString
			wishes           sql.NullString
			AvatarURL        sql.NullString

			winner                   models.Organization
			winnerID                 sql.NullInt64
			winnerBrandName          sql.NullString
			winnerFullName           sql.NullString
			winnerShortName          sql.NullString
			winnerINN                sql.NullString
			winnerOKPO               sql.NullString
			winnerOGRN               sql.NullString
			winnerKPP                sql.NullString
			winnerTaxCode            sql.NullString
			winnerAddress            sql.NullString
			winnerAvatarURL          sql.NullString
			winnerVerificationStatus sql.NullInt16
			winnerIsContractor       sql.NullBool
			winnerIsBanned           sql.NullBool
			winnerCreatedAt          sql.NullTime
			winnerUpdatedAt          sql.NullTime
		)

		err = rows.Scan(
			&tender.ID,
			&tender.City.ID,
			&tenderServiceIDs,
			&tenderObjectIDs,
			&tender.Name,
			&tender.Price,
			&tender.IsContractPrice,
			&tender.IsNDSPrice,
			&tender.FloorSpace,
			&description,
			&wishes,
			&tender.Specification,
			pq.Array(&tender.Attachments),
			&tender.Status,
			&tender.VerificationStatus,
			&tender.IsDraft,
			&tender.ReceptionStart,
			&tender.ReceptionEnd,
			&tender.WorkStart,
			&tender.WorkEnd,
			&tender.CreatedAt,
			&tender.UpdatedAt,
			&tender.City.Name,
			&tender.City.ID,
			&tender.City.Region.Name,
			&tender.City.Region.ID,
			&tender.Organization.ID,
			&tender.Organization.BrandName,
			&tender.Organization.FullName,
			&tender.Organization.ShortName,
			&tender.Organization.INN,
			&tender.Organization.OKPO,
			&tender.Organization.OGRN,
			&tender.Organization.KPP,
			&tender.Organization.TaxCode,
			&tender.Organization.Address,
			&AvatarURL,
			&tender.Organization.Emails,
			&tender.Organization.Phones,
			&tender.Organization.Messengers,
			&tender.Organization.VerificationStatus,
			&tender.Organization.IsContractor,
			&tender.Organization.IsBanned,
			&tender.Organization.CustomerInfo,
			&tender.Organization.ContractorInfo,
			&tender.Organization.CreatedAt,
			&tender.Organization.UpdatedAt,
			&winnerID,
			&winnerBrandName,
			&winnerFullName,
			&winnerShortName,
			&winnerINN,
			&winnerOKPO,
			&winnerOGRN,
			&winnerKPP,
			&winnerTaxCode,
			&winnerAddress,
			&winnerAvatarURL,
			&winner.Emails,
			&winner.Phones,
			&winner.Messengers,
			&winnerVerificationStatus,
			&winnerIsContractor,
			&winnerIsBanned,
			&winner.CustomerInfo,
			&winner.ContractorInfo,
			&winnerCreatedAt,
			&winnerUpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		tender.Organization.AvatarURL = AvatarURL.String
		tender.Description = description.String
		tender.Wishes = wishes.String

		if winnerID.Valid {
			tender.WinnerOrganization = models.Optional[models.Organization]{
				Value: models.Organization{
					ID:                 int(winnerID.Int64),
					BrandName:          winnerBrandName.String,
					FullName:           winnerFullName.String,
					ShortName:          winnerShortName.String,
					INN:                winnerINN.String,
					OKPO:               winnerOKPO.String,
					OGRN:               winnerOGRN.String,
					KPP:                winnerKPP.String,
					TaxCode:            winnerTaxCode.String,
					Address:            winnerAddress.String,
					AvatarURL:          winnerAvatarURL.String,
					VerificationStatus: models.VerificationStatus(winnerVerificationStatus.Int16),
					IsContractor:       winnerIsContractor.Bool,
					IsBanned:           winnerIsBanned.Bool,
					Emails:             winner.Emails,
					Phones:             winner.Phones,
					Messengers:         winner.Messengers,
					CustomerInfo:       winner.CustomerInfo,
					ContractorInfo:     winner.ContractorInfo,
					CreatedAt:          winnerCreatedAt.Time,
					UpdatedAt:          winnerUpdatedAt.Time,
				},
				Set: true,
			}
		}

		tenderServiceIDsConverted := convert.Slice[[]int64, []int](tenderServiceIDs, func(i int64) int { return int(i) })
		tenderObjectIDsConverted := convert.Slice[[]int64, []int](tenderObjectIDs, func(i int64) int { return int(i) })

		tenders = append(tenders, tender)
		tenderToServices[tender.ID] = tenderServiceIDsConverted
		tenderToObjects[tender.ID] = tenderObjectIDsConverted
		serviceIDs = append(serviceIDs, tenderServiceIDsConverted...)
		objectIDs = append(objectIDs, tenderObjectIDsConverted...)
	}

	services, err := s.catalogStore.GetServicesByIDs(ctx, qe, deduplicate.Deduplicate(serviceIDs))
	if err != nil {
		return nil, fmt.Errorf("get services: %w", err)
	}

	objects, err := s.catalogStore.GetObjectsByIDs(ctx, qe, deduplicate.Deduplicate(objectIDs))
	if err != nil {
		return nil, fmt.Errorf("get objects: %w", err)
	}

	for i, tender := range tenders {
		tenderServiceIDs := tenderToServices[tender.ID]
		tenderServices := make([]models.Service, 0, len(serviceIDs))

		for _, id := range tenderServiceIDs {
			tenderServices = append(tenderServices, services[id])
		}

		tenderObjectIDs := tenderToObjects[tender.ID]
		tenderObjects := make([]models.Object, 0, len(objectIDs))

		for _, id := range tenderObjectIDs {
			tenderObjects = append(tenderObjects, objects[id])
		}

		tenders[i].Services = tenderServices
		tenders[i].Objects = tenderObjects
	}
	
	return tenders, nil
}
