package verification

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) GetTendersRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	builder := squirrel.Select(
		"vr.id",
		"vr.reviewer_user_id",
		"u.email",
		"u.phone",
		"u.first_name",
		"u.last_name",
		"u.middle_name",
		"u.avatar_url",
		"u.email_verified",
		"u.is_banned",
		"u.created_at AS user_created_at",
		"u.updated_at AS user_updated_at",
		"e.position",
		"e.role",
		"vr.object_type",
		"vr.content",
		"vr.attachments",
		"vr.status",
		"vr.created_at AS vr_created_at",
		"vr.reviewed_at AS vr_reviewed_at",
		"t.id").
		From("verification_requests vr").
		Join("users u ON vr.reviewer_user_id = u.id").
		Join("employee e ON e.user_id = u.id").
		Join("tenders t ON vr.object_id = t.id").
		Where(
			squirrel.Eq{
				"vr.object_type": params.ObjectType,
				"vr.status":      params.Status},
		).
		Offset(params.Offset).
		Limit(params.Limit).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var requests []models.VerificationRequest[models.VerificationObject]

	for rows.Next() {
		var (
			request       models.VerificationRequest[models.VerificationObject]
			userAvatarURL sql.NullString
			tenderID      int
		)

		if err := rows.Scan(
			&request.ID,
			&request.Reviewer.ID,
			&request.Reviewer.Email,
			&request.Reviewer.Phone,
			&request.Reviewer.FirstName,
			&request.Reviewer.LastName,
			&request.Reviewer.MiddleName,
			&userAvatarURL,
			&request.Reviewer.EmailVerified,
			&request.Reviewer.IsBanned,
			&request.Reviewer.CreatedAt,
			&request.Reviewer.UpdatedAt,
			&request.Reviewer.Position,
			&request.Reviewer.Role,
			&request.ObjectType,
			&request.Content,
			pq.Array(&request.Attachments),
			&request.Status,
			&request.CreatedAt,
			&request.ReviewedAt,
			&tenderID); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		request.ObjectID = tenderID
		request.Object = &models.Tender{}

		requests = append(requests, request)
	}

	if len(requests) == 0 {
		return nil, fmt.Errorf("requests not found")
	}

	return requests, nil
}

func (s *VerificationStore) GetOrganizationRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	builder := squirrel.Select(
		"vr.id",
		"vr.reviewer_user_id",
		"u.email",
		"u.phone",
		"u.first_name",
		"u.last_name",
		"u.middle_name",
		"u.avatar_url",
		"u.email_verified",
		"u.is_banned",
		"u.created_at AS user_created_at",
		"u.updated_at AS user_updated_at",
		"e.position",
		"e.role",
		"vr.object_type",
		"vr.content",
		"vr.attachments",
		"vr.status",
		"vr.created_at AS vr_created_at",
		"vr.reviewed_at AS vr_reviewed_at",
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
	).From("verification_requests vr").
		Join("users u ON vr.reviewer_user_id = u.id").
		Join("employee e ON e.user_id = u.id").
		Join("organizations o ON vr.object_id = o.id").
		Where(
			squirrel.Eq{
				"vr.object_type": params.ObjectType,
				"vr.status":      params.Status},
		).
		Offset(params.Offset).
		Limit(params.Limit).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var requests []models.VerificationRequest[models.VerificationObject]

	for rows.Next() {
		var (
			request       models.VerificationRequest[models.VerificationObject]
			userAvatarURL sql.NullString
			orgAvatarURL  sql.NullString
			org           models.Organization
		)

		if err := rows.Scan(
			&request.ID,
			&request.Reviewer.ID,
			&request.Reviewer.Email,
			&request.Reviewer.Phone,
			&request.Reviewer.FirstName,
			&request.Reviewer.LastName,
			&request.Reviewer.MiddleName,
			&userAvatarURL,
			&request.Reviewer.EmailVerified,
			&request.Reviewer.IsBanned,
			&request.Reviewer.CreatedAt,
			&request.Reviewer.UpdatedAt,
			&request.Reviewer.Position,
			&request.Reviewer.Role,
			&request.ObjectType,
			&request.Content,
			pq.Array(&request.Attachments),
			&request.Status,
			&request.CreatedAt,
			&request.ReviewedAt,
			&org.ID,
			&org.BrandName,
			&org.FullName,
			&org.ShortName,
			&org.INN,
			&org.OKPO,
			&org.OGRN,
			&org.KPP,
			&org.TaxCode,
			&org.Address,
			&orgAvatarURL,
			&org.Emails,
			&org.Phones,
			&org.Messengers,
			&org.VerificationStatus,
			&org.IsContractor,
			&org.IsBanned,
			&org.CustomerInfo,
			&org.ContractorInfo,
			&org.CreatedAt,
			&org.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		request.Reviewer.AvatarURL = userAvatarURL.String
		org.AvatarURL = orgAvatarURL.String
		request.Object = org

		requests = append(requests, request)
	}

	if len(requests) == 0 {
		return nil, fmt.Errorf("requests not found")
	}

	return requests, nil
}

func (s *VerificationStore) GetCommentRequests(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	builder := squirrel.Select(
		"vr.id",
		"vr.reviewer_user_id",
		"u.email",
		"u.phone",
		"u.first_name",
		"u.last_name",
		"u.middle_name",
		"u.avatar_url",
		"u.email_verified",
		"u.is_banned",
		"u.created_at AS user_created_at",
		"u.updated_at AS user_updated_at",
		"e.position",
		"e.role",
		"vr.object_type",
		"vr.content",
		"vr.attachments",
		"vr.status",
		"vr.created_at AS vr_created_at",
		"vr.reviewed_at AS vr_reviewed_at",
		"c.id",
		"c.object_type",
		"c.object_id",
		"c.content",
		"c.attachments",
		"c.verification_status",
		"c.created_at",
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
	).From("verification_requests vr").
		Join("users u ON vr.reviewer_user_id = u.id").
		Join("employee e ON e.user_id = u.id").
		Join("comments c ON vr.object_id = c.id").
		Join("organizations o ON c.organization_id = o.id").
		Where(
			squirrel.Eq{
				"vr.object_type": params.ObjectType,
				"vr.status":      params.Status},
		).
		Offset(params.Offset).
		Limit(params.Limit).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var requests []models.VerificationRequest[models.VerificationObject]

	for rows.Next() {
		var (
			request       models.VerificationRequest[models.VerificationObject]
			userAvatarURL sql.NullString
			orgAvatarURL  sql.NullString
			comment       models.Comment
		)

		if err := rows.Scan(
			&request.ID,
			&request.Reviewer.ID,
			&request.Reviewer.Email,
			&request.Reviewer.Phone,
			&request.Reviewer.FirstName,
			&request.Reviewer.LastName,
			&request.Reviewer.MiddleName,
			&userAvatarURL,
			&request.Reviewer.EmailVerified,
			&request.Reviewer.IsBanned,
			&request.Reviewer.CreatedAt,
			&request.Reviewer.UpdatedAt,
			&request.Reviewer.Position,
			&request.Reviewer.Role,
			&request.ObjectType,
			&request.Content,
			pq.Array(&request.Attachments),
			&request.Status,
			&request.CreatedAt,
			&request.ReviewedAt,
			&comment.ID,
			&comment.ObjectType,
			&comment.ObjectID,
			&comment.Content,
			pq.Array(&comment.Attachments),
			&comment.VerificationStatus,
			&comment.CreatedAt,
			&comment.Organization.ID,
			&comment.Organization.BrandName,
			&comment.Organization.FullName,
			&comment.Organization.ShortName,
			&comment.Organization.INN,
			&comment.Organization.OKPO,
			&comment.Organization.OGRN,
			&comment.Organization.KPP,
			&comment.Organization.TaxCode,
			&comment.Organization.Address,
			&orgAvatarURL,
			&comment.Organization.Emails,
			&comment.Organization.Phones,
			&comment.Organization.Messengers,
			&comment.Organization.VerificationStatus,
			&comment.Organization.IsContractor,
			&comment.Organization.IsBanned,
			&comment.Organization.CustomerInfo,
			&comment.Organization.ContractorInfo,
			&comment.Organization.CreatedAt,
			&comment.Organization.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		request.Reviewer.AvatarURL = userAvatarURL.String
		comment.Organization.AvatarURL = orgAvatarURL.String
		request.Object = comment

		requests = append(requests, request)
	}

	if len(requests) == 0 {
		return nil, fmt.Errorf("requests not found")
	}

	return requests, nil
}
