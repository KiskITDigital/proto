package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) CreateComment(ctx context.Context, params service.CommentCreateParams) error {
	organizationID := contextor.GetOrganizationID(ctx)

	err := s.tenderStore.CreateComment(ctx, s.psql.DB(), store.CommentCreateParams{
		ObjectType:     models.ObjectTypeTender,
		ObjectID:       params.TenderID,
		OrganizationID: organizationID,
		Content:        params.Content,
		Attachments:    params.Attachments,
	})
	if err != nil {
		return fmt.Errorf("creating comment %w", err)
	}

	return nil
}
