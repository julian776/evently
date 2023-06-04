package models

type Reminder struct {
	EventId        string
	DateToSend     string
	MessageToSend  string
	EmailsToNotify []string
}
