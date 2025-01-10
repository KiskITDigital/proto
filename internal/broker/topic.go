package broker

type Topic string

const (
	UbratoUserRegisteredSubject  = "ubrato.user.registered"
	UbratoSurveySubmittedSubject = "ubrato.survey.submitted"
	UbratoUserConfirmEmail       = "email.send.confirmation"
	UbratoUserEmailResetPass     = "email.send.resetpass"

	UbratoNotificationUsers        = "notification.users.send"
	UbratoUserRegistration         = "user.registration"
	UbratoOrganizationVerification = "organization.verification"

	UbratoTenderVerification               = "tender.verification"
	UbratoTenderAdditionVerification       = "tender.addition.verification"
	UbratoTenderInvitation                 = "tender.invitation"
	UbratoTenderQuestionAnswerVerification = "tender.question.answer.verification"
)
