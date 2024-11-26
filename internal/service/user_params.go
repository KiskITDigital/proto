package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type UserCreateEmployeeParams struct {
	Email      string
	Phone      string
	Password   string
	FirstName  string
	LastName   string
	MiddleName models.Optional[string]
	Role       models.UserRole
	Position   string
}
