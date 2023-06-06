package broker

import (
	"context"
	"notifier/domain/broker/models"
)

type BrokerListener interface {
	AddQueueToListen(queueName string) error
	AddMessageHandler(typeMessage string, handler models.HandlerFunc)
	Listen(ctx context.Context)
	Stop()
}
