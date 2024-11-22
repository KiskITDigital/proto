package comment

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CommentStore) GetComments(ctx context.Context, qe store.QueryExecutor, params store.CommentGetParams) ([]models.Comment, error) {
	builder := squirrel.
		Select(
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
			"o.is_contractor",
			"o.is_banned",
			"o.created_at",
			"o.updated_at",
		).
		From("comments as c").
		Join("organizations AS o ON o.id = c.organization_id").
		Where(squirrel.Eq{"object_type": params.ObjectType}).
		Where(squirrel.Eq{"object_id": params.ObjectID}).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.Attachments,
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
			&comment.Organization.AvatarURL,
			&comment.Organization.Emails,
			&comment.Organization.Phones,
			&comment.Organization.Messengers,
			&comment.Organization.IsContractor,
			&comment.Organization.IsBanned,
			&comment.Organization.CreatedAt,
			&comment.Organization.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return comments, nil
}
