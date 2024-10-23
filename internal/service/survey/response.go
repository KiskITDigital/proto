package survey

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	commandsv1 "gitlab.ubrato.ru/ubrato/core/internal/gen/amo-sync-pb/commands/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/gen/amo-sync-pb/models/v1"
	"google.golang.org/protobuf/proto"
)

type ResponseParams struct {
	Name     string
	Type     string
	Phone    string
	Question string
}

func (s *Service) Response(ctx context.Context, params ResponseParams) error {
	b, err := proto.Marshal(&commandsv1.CreateLead{
		Name:     params.Name,
		Status:   modelsv1.Status_IncomingRequestsNewLead,
		Pipeline: modelsv1.Pipeline_IncomingRequests,
		Contact: &commandsv1.CreateLead_ContactModel{
			ContactModel: &modelsv1.Contact{
				FirstName:  params.Name,
				LastName:   "",
				MiddleName: "",
				Email:      "",
				Phone:      params.Phone,
			},
		},
		InitialMessage: &params.Question,
	})
	if err != nil {
		return fmt.Errorf("marshal proto: %w", err)
	}

	return s.broker.Publish(ctx, broker.AmoCreateLeadTopic, b)
}
