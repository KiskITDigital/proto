package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type VerificationRequestUpdateStatusParams struct {
	RequesID int
	Status   models.VerificationStatus
}

type VerificationObjectUpdateStatusParams struct {
	Object models.VerificationObject
	Status models.VerificationStatus
}

type VerificationRequestsObjectGetParams struct {
	ObjectType models.ObjectType
	Status     []models.VerificationStatus
	Offset     uint64
	Limit      uint64
}
