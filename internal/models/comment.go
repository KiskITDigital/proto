package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type Comment struct {
	VerificationObject

	ID                 int
	ObjectType         string
	ObjectID           int
	Content            string
	Attachments        []string
	VerificationStatus VerificationStatus
	CreatedAt          time.Time
}

func ConvertCommentModelToApi(comment Comment) api.Comment {
	return api.Comment{
		ID:          comment.ID,
		Content:     comment.Content,
		Attachments: comment.Attachments,
		CreatedAt:   comment.CreatedAt,
	}
}

func (c Comment) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:    api.CommentVerificationRequestObject,
		Comment: ConvertCommentModelToApi(c),
	}
}
