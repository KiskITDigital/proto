package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestCreateParams struct {
	ObjectID    int
	ObjectType  models.ObjectType
	Attachments models.Attachments
}

type VerificationRequestUpdateStatusParams struct {
	UserID        int
	RequestID     int
	Status        models.VerificationStatus
	ReviewComment models.Optional[string]
}

type VerificationObjectUpdateStatusResult struct {
	ObjectID   int
	ObjectType models.ObjectType
}

type VerificationRequestsObjectGetParams struct {
	ObjectID   models.Optional[int]
	ObjectType models.Optional[models.ObjectType]
	Status     []models.VerificationStatus
	Offset     uint64
	Limit      uint64
}
