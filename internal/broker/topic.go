package broker

type Topic string

const (
	UbratoUserRegisteredSubject  = "ubrato.user.registered"
	UbratoSurveySubmittedSubject = "ubrato.survey.submitted"
	UbratoUserConfirmEmail       = "email.send.confirmation"
	UbratoUserEmailResetPass     = "email.send.resetpass"
	UbratoNotificationUsers      = "notification.users.send"
)