package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type ObjectType int

const (
	ObjectTypeOrganization ObjectType = iota
	ObjectTypeTender
	ObjectTypeComment
)

var mapVerificationObjectType = map[ObjectType]api.ObjectType{
	ObjectTypeOrganization: api.ObjectTypeOrganization,
	ObjectTypeTender:       api.ObjectTypeTender,
	ObjectTypeComment:      api.ObjectTypeComment,
}

func (v ObjectType) ToAPI() api.ObjectType {
	return mapVerificationObjectType[v]
}
