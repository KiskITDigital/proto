package errstore

import "errors"

const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

var (
	ErrQuestionnaireExist = errors.New("the questionnaire has been completed")
)

var (
	ErrUserNotFound = errors.New("user not found")
)