package errstore

import "errors"

const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

var (
	ErrQuestionnaireExist    = errors.New("questionnaire has been completed")
	ErrQuestionnaireNotFound = errors.New("questionnaire not found")
)

var (
	ErrUserNotFound = errors.New("user not found")
)

var (
	ErrOrganizationNotFound = errors.New("organization not found")
)

var (
	ErrQuestionAnswerUniqueViolation = errors.New("answer to question already exists")
	ErrQuestionAnswerFKViolation     = errors.New("foreign key violation on question_answer")
)
