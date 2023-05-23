package events

import "events-manager/domain/broker"

func LoadEventsService(
	brokerPublisher broker.BrokerPublisher,
	settings EventsSettings,
) {
	brokerPublisher.QueueDeclare(settings.Queue)
}
