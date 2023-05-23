package events

type EventsSettings struct {
	Queue string `envconfig:"EVENTS_QUEUE" default:"events"`
}
