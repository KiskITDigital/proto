package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type VerificationObject interface {
	ToVerificationObject() api.VerificationRequestObject
}

type VerificationRequest[T VerificationObject] struct {
	ID            int
	Reviewer      EmployeeUser
	ObjectType    ObjectType
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
		Reviewer:      api.OptEmployeeUser{Value: ConvertEmployeeUserModelToApi(request.Reviewer), Set: request.Reviewer.ID != 0},
		ObjectType:    api.ObjectType(request.ObjectType.ToAPI()),
		Object:        request.Object.ToVerificationObject(),
		Content:       request.Content,
		Attachments:   request.Attachments,
		Status:        request.Status.ToAPI(),
		ReviewComment: api.OptString{Value: request.ReviewComment, Set: request.ReviewComment != ""},
		CreatedAt:     request.CreatedAt,
		ReviewedAt:    api.OptDateTime{Value: request.ReviewedAt, Set: !request.ReviewedAt.IsZero()},
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

var mapApiToVerificationStatus = map[api.VerificationStatus]VerificationStatus{
	api.VerificationStatusUnverified: VerificationStatusUnverified,
	api.VerificationStatusInReview:   VerificationStatusInReview,
	api.VerificationStatusDeclined:   VerificationStatusDeclined,
	api.VerificationStatusApproved:   VerificationStatusApproved,
}

func APIToVerificationStatus(apiStatus api.VerificationStatus) VerificationStatus {
	status, ok := mapApiToVerificationStatus[apiStatus]
	if !ok {
		return VerificationStatusUnverified
	}
	return status
}
