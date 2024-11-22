package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type CommentCreateParams struct {
	ObjectType     models.ObjectType
	ObjectID       int
	OrganizationID int
	Content        string
	Attachments    []string
}
