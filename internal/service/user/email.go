package user

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
)

func (s *Service) ReqEmailVerification(ctx context.Context, email string) error {
	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{Email: email})
	if err != nil {
		return fmt.Errorf("store get user: %v", err)
	}

	if len(users) == 0 {
		cerr.Wrap(
			errors.New("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %s email not found", email),
			nil,
		)
	}

	user := users[0]

	code, err := crypto.GenerateTOTPCode(user.TOTPSalt)
	if err != nil {
		return fmt.Errorf("generate topt: %v", err)
	}

	confirmPb, err := proto.Marshal(&modelsv1.EmailConfirmation{
		Email: user.Email,
		Salt:  code,
	})
	if err != nil {
		return fmt.Errorf("marshal proto: %w", err)
	}

	return s.broker.Publish(ctx, broker.UbratoUserConfirmEmail, confirmPb)
}
