package broker

import (
	"context"
)

type BrokerListener interface {
	AddQueueToListen(queueName string) error
	AddMessageHandler(typeMessage string, handler func(ctx context.Context, command string) interface{})
	PublishMessageWithContext(
		ctx context.Context,
		queueName string,
		message interface{},
		typeMessage string,
	) error
	Listen() error
	Stop()
	//ListenMessages(ctx context.Context, appName string, brokerId string) chan<- models.Message
}
