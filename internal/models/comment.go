package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type Comment struct {
	VerificationObject

	ID                 int
	Organization       Organization
	ObjectType         ObjectType
	ObjectID           int
	Title              string
	Content            string
	Attachments        []string
	VerificationStatus VerificationStatus
	CreatedAt          time.Time
}

func ConvertCommentModelToApi(comment Comment) api.Comment {
	return api.Comment{
		ID:                 comment.ID,
		Organization:       ConvertOrganizationModelToApi(comment.Organization),
		Title:              comment.Title,
		Content:            comment.Content,
		Attachments:        comment.Attachments,
		VerificationStatus: string(comment.VerificationStatus.ToAPI()),
		CreatedAt:          comment.CreatedAt,
	}
}

func (c Comment) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:    api.CommentVerificationRequestObject,
		Comment: ConvertCommentModelToApi(c),
	}
}
