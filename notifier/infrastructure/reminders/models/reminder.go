package models

import "time"

type Reminder struct {
	EventId        string    `bson:"eventId,omitempty"`
	DateToSend     time.Time `bson:"dateToSend,omitempty"`
	MessageToSend  string    `bson:"messageToSend,omitempty"`
	EmailsToNotify []string  `bson:"emailsToNotify,omitempty"`
}
