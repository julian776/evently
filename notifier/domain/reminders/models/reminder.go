package models

type Reminder struct {
	EventId        string   `json:"eventId,omitempty"`
	DateToSend     string   `json:"dateToSend,omitempty"`
	MessageToSend  string   `json:"messageToSend,omitempty"`
	EmailsToNotify []string `json:"emailsToNotify,omitempty"`
}
