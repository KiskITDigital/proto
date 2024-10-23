package broker

type Topic string

const (
	AmoCreateCompanyTopic Topic = "jobs.amo.create_company"
	AmoCreateLeadTopic    Topic = "jobs.amo.create_lead"
)
