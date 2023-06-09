package models

import "time"

type Reminder struct {
	EventId        string    `json:"eventId,omitempty"`
	DateToSend     time.Time `json:"dateToSend,omitempty"`
	MessageToSend  string    `json:"messageToSend,omitempty"`
	EmailsToNotify []string  `json:"emailsToNotify,omitempty"`
}
