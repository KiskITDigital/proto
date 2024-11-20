package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type VerificationObject interface {
	ToVerificationObject() api.VerificationRequestObject
}

type VerificationRequest[T any] struct {
	ID            int
	Reviewer      EmployeeUser
	ObjectType    string
	ObjectID      int
	Object        T
	Content       string
	Attachments   []string
	Status        VerificationStatus
	ReviewComment string
	CreatedAt     time.Time
	ReviewedAt    time.Time
}

func VerificationRequestModelToApi[T VerificationObject](request VerificationRequest[T]) api.VerificationRequest {
	return api.VerificationRequest{
		ID:            request.ID,
		Reviewer:      api.NewOptEmployeeUser(ConvertEmployeeUserModelToApi(request.Reviewer)),
		ObjectType:    api.ObjectType(request.ObjectType),
		Object:        request.Object.ToVerificationObject(),
		Attachments:   request.Attachments,
		Status:        request.Status.ToAPI(),
		ReviewComment: api.NewOptString(request.ReviewComment),
		CreatedAt:     request.CreatedAt,
		ReviewedAt:    api.NewOptDateTime(request.ReviewedAt),
	}
}

type VerificationStatus int

const (
	VerificationStatusUnverified VerificationStatus = iota
	VerificationStatusInReview
	VerificationStatusDeclined
	VerificationStatusApproved
)

var mapVerificationStatus = map[VerificationStatus]api.VerificationStatus{
	VerificationStatusUnverified: api.VerificationStatusUnverified,
	VerificationStatusInReview:   api.VerificationStatusInReview,
	VerificationStatusDeclined:   api.VerificationStatusDeclined,
	VerificationStatusApproved:   api.VerificationStatusApproved,
}

func (v VerificationStatus) ToAPI() api.VerificationStatus {
	return mapVerificationStatus[v]
}
