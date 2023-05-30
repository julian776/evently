package users

import "events-manager/domain/broker"

func LoadEventsService(
	brokerPublisher broker.BrokerPublisher,
	settings UsersSettings,
) {
	brokerPublisher.QueueDeclare(settings.Queue)
}
