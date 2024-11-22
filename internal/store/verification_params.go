package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestUpdateStatusParams struct {
	ID     int
	Status models.VerificationStatus
}

type VerificationObjectUpdateStatusResult struct {
	ObjectID   int
	ObjectType models.VerificationObjectType
}
