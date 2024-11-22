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
