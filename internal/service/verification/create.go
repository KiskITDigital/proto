package verification

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.VerificationRequestCreateParams) error {
	return s.verificationStore.Create(ctx, s.psql.DB(), store.VerificationRequestCreateParams{
		ObjectID:    params.ObjectID,
		ObjectType:  params.ObjectType,
		Attachments: params.Attachments,
	})
}
