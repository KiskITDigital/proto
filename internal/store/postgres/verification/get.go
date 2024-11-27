package verification

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) GetByIDWithEmptyObject(ctx context.Context, qe store.QueryExecutor, requestID int) (models.VerificationRequest[models.VerificationObject], error) {
	requests, err := s.GetWithEmptyObject(ctx, qe, store.VerificationRequestsObjectGetParams{
		ObjectID: models.Optional[int]{Value: requestID, Set: true},
		Limit:    1,
	})
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get objects: %w", err)
	}

	if len(requests) == 0 {
		return models.VerificationRequest[models.VerificationObject]{}, errors.New("object not found")
	}

	return requests[0], nil
}

func (s *VerificationStore) GetWithEmptyObject(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
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
		"vr.object_id",
		"vr.content",
		"vr.attachments",
		"vr.status",
		"vr.review_comment",
		"vr.created_at AS vr_created_at",
		"vr.reviewed_at AS vr_reviewed_at").
		From("verification_requests vr").
		LeftJoin("users u ON vr.reviewer_user_id = u.id").
		LeftJoin("employee e ON e.user_id = u.id").
		Offset(params.Offset).
		Limit(params.Limit).
		PlaceholderFormat(squirrel.Dollar)

	if params.ObjectType.Set {
		builder = builder.Where(squirrel.Eq{"vr.object_type": params.ObjectType.Value})
	}

	if params.ObjectID.Set {
		builder = builder.Where(squirrel.Eq{"vr.id": params.ObjectID.Value})
	}

	if len(params.Status) != 0 {
		builder = builder.Where(squirrel.Eq{"vr.status": params.Status})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var requests []models.VerificationRequest[models.VerificationObject]

	for rows.Next() {
		var (
			request models.VerificationRequest[models.VerificationObject]

			reviewerID            sql.NullInt64
			reviewerEmail         sql.NullString
			reviewerPhone         sql.NullString
			reviewerFirstName     sql.NullString
			reviewerLastName      sql.NullString
			reviewerMiddleName    sql.NullString
			reviewerAvatarURL     sql.NullString
			reviewerEmailVerified sql.NullBool
			reviewerIsBanned      sql.NullBool
			reviewerCreatedAt     sql.NullTime
			reviewerUpdatedAt     sql.NullTime
			reviewerPosition      sql.NullString
			reviewerRole          sql.NullInt16

			requestContent         sql.NullString
			requestReviewerComment sql.NullString
			requestReviewedAt      sql.NullTime
		)

		if err := rows.Scan(
			&request.ID,
			&reviewerID,
			&reviewerEmail,
			&reviewerPhone,
			&reviewerFirstName,
			&reviewerLastName,
			&reviewerMiddleName,
			&reviewerAvatarURL,
			&reviewerEmailVerified,
			&reviewerIsBanned,
			&reviewerCreatedAt,
			&reviewerUpdatedAt,
			&reviewerPosition,
			&reviewerRole,
			&request.ObjectType,
			&request.ObjectID,
			&requestContent,
			pq.Array(&request.Attachments),
			&request.Status,
			&requestReviewerComment,
			&request.CreatedAt,
			&requestReviewedAt); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		request.Content = requestContent.String
		request.ReviewComment = requestReviewerComment.String
		request.ReviewedAt = requestReviewedAt.Time

		if reviewerID.Valid {
			request.Reviewer = models.EmployeeUser{
				User: models.User{
					ID:            int(reviewerID.Int64),
					Email:         reviewerEmail.String,
					Phone:         reviewerPhone.String,
					FirstName:     reviewerFirstName.String,
					LastName:      reviewerLastName.String,
					MiddleName:    models.Optional[string]{Value: reviewerMiddleName.String, Set: reviewerMiddleName.Valid},
					AvatarURL:     reviewerAvatarURL.String,
					EmailVerified: reviewerEmailVerified.Bool,
					IsBanned:      reviewerIsBanned.Bool,
					CreatedAt:     reviewerCreatedAt.Time,
					UpdatedAt:     reviewerUpdatedAt.Time,
				},
				Position: reviewerPosition.String,
				Role:     models.UserRole(reviewerRole.Int16),
			}
		}

		requests = append(requests, request)
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
		"vr.review_comment",
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
		LeftJoin("users u ON vr.reviewer_user_id = u.id").
		LeftJoin("employee e ON e.user_id = u.id").
		Join("organizations o ON vr.object_id = o.id").
		Where(
			squirrel.Eq{
				"vr.object_type": models.ObjectTypeOrganization,
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
			request models.VerificationRequest[models.VerificationObject]

			reviewerID            sql.NullInt64
			reviewerEmail         sql.NullString
			reviewerPhone         sql.NullString
			reviewerFirstName     sql.NullString
			reviewerLastName      sql.NullString
			reviewerMiddleName    sql.NullString
			reviewerAvatarURL     sql.NullString
			reviewerEmailVerified sql.NullBool
			reviewerIsBanned      sql.NullBool
			reviewerCreatedAt     sql.NullTime
			reviewerUpdatedAt     sql.NullTime
			reviewerPosition      sql.NullString
			reviewerRole          sql.NullInt16

			requestContent         sql.NullString
			requestReviewerComment sql.NullString
			requestReviewedAt      sql.NullTime

			org          models.Organization
			orgAvatarURL sql.NullString
		)

		if err := rows.Scan(
			&request.ID,
			&reviewerID,
			&reviewerEmail,
			&reviewerPhone,
			&reviewerFirstName,
			&reviewerLastName,
			&reviewerMiddleName,
			&reviewerAvatarURL,
			&reviewerEmailVerified,
			&reviewerIsBanned,
			&reviewerCreatedAt,
			&reviewerUpdatedAt,
			&reviewerPosition,
			&reviewerRole,
			&request.ObjectType,
			&requestContent,
			pq.Array(&request.Attachments),
			&request.Status,
			&requestReviewerComment,
			&request.CreatedAt,
			&requestReviewedAt,
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

		org.AvatarURL = orgAvatarURL.String
		request.Object = org

		request.Content = requestContent.String
		request.ReviewComment = requestReviewerComment.String
		request.ReviewedAt = requestReviewedAt.Time

		if reviewerID.Valid {
			request.Reviewer = models.EmployeeUser{
				User: models.User{
					ID:            int(reviewerID.Int64),
					Email:         reviewerEmail.String,
					Phone:         reviewerPhone.String,
					FirstName:     reviewerFirstName.String,
					LastName:      reviewerLastName.String,
					MiddleName:    models.Optional[string]{Value: reviewerMiddleName.String, Set: reviewerMiddleName.Valid},
					AvatarURL:     reviewerAvatarURL.String,
					EmailVerified: reviewerEmailVerified.Bool,
					IsBanned:      reviewerIsBanned.Bool,
					CreatedAt:     reviewerCreatedAt.Time,
					UpdatedAt:     reviewerUpdatedAt.Time,
				},
				Position: reviewerPosition.String,
				Role:     models.UserRole(reviewerRole.Int16),
			}
		}

		requests = append(requests, request)
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
		"c.title",
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
		Where(squirrel.Eq{"vr.object_type": models.ObjectTypeComment}).
		Offset(params.Offset).
		Limit(params.Limit).
		PlaceholderFormat(squirrel.Dollar)

	if params.ObjectID.Set {
		builder = builder.Where(squirrel.Eq{"vr.id": params.ObjectID.Value})
	}

	if len(params.Status) != 0 {
		builder = builder.Where(squirrel.Eq{"vr.status": params.Status})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	var requests []models.VerificationRequest[models.VerificationObject]

	for rows.Next() {
		var (
			request        models.VerificationRequest[models.VerificationObject]
			userAvatarURL  sql.NullString
			userMiddleName sql.NullString
			orgAvatarURL   sql.NullString
			comment        models.Comment
		)

		if err := rows.Scan(
			&request.ID,
			&request.Reviewer.ID,
			&request.Reviewer.Email,
			&request.Reviewer.Phone,
			&request.Reviewer.FirstName,
			&request.Reviewer.LastName,
			&userMiddleName,
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
			&comment.Title,
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

		comment.Organization.AvatarURL = orgAvatarURL.String
		request.Object = comment

		request.Reviewer.AvatarURL = userAvatarURL.String
		request.Reviewer.MiddleName = models.Optional[string]{Value: userMiddleName.String, Set: userMiddleName.Valid}

		requests = append(requests, request)
	}

	return requests, nil
}
