package broker

import (
	"context"
)

type BrokerListener interface {
	AddQueueToListen(queueName string) error
	AddMessageHandler(queueName string, typeMessage string, handler func(ctx context.Context, command string) interface{})
	Listen() error
	Stop()
}
