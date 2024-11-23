package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestCreateParams struct {
	ObjectID    int
	ObjectType  models.ObjectType
	Attachments []string
}

type VerificationRequestUpdateStatusParams struct {
	ID     int
	Status models.VerificationStatus
}

type VerificationObjectUpdateStatusResult struct {
	ObjectID   int
	ObjectType models.ObjectType
}
