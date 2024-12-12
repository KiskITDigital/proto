package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestUpdateStatusParams struct {
	UserID        int
	RequesID      int
	Status        models.VerificationStatus
	ReviewComment models.Optional[string]
}

type VerificationObjectUpdateStatusParams struct {
	Object models.VerificationObject
	Status models.VerificationStatus
}

type VerificationRequestsObjectGetParams struct {
	ObjectType models.ObjectType
	ObjectID   models.Optional[int]
	Status     []models.VerificationStatus
	Offset     uint64
	Limit      uint64
}
