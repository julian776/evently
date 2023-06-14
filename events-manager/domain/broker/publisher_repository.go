package broker

import (
	"context"
)

//go:generate mockery --name=BrokerPublisher
type BrokerPublisher interface {
	QueueDeclare(queueName string) error
	PublishMessageWithContext(
		ctx context.Context,
		queueName string,
		message map[string]any,
		typeMessage string,
	) error
	Stop()
	//ListenMessages(ctx context.Context, appName string, brokerId string) chan<- models.Message
}
