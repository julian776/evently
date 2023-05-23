package broker

import (
	"context"
)

type BrokerPublisher interface {
	QueueDeclare(queueName string) error
	PublishMessageWithContext(
		ctx context.Context,
		queueName string,
		message interface{},
		typeMessage string,
	) error
	Stop()
	//ListenMessages(ctx context.Context, appName string, brokerId string) chan<- models.Message
}
