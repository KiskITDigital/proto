package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func(s *Service) CreateVerificationRequest(ctx context.Context, params service.OrganizationCreateVerificationRequestParams) error {
	return s.verificationStore.Create(ctx, s.psql.DB(), store.VerificationRequestCreateParams{
		ObjectID: params.OrganizationID,
		ObjectType: models.ObjectTypeOrganization,
		Attachments: params.Attachments,
	})
}