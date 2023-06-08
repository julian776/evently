package listener

import (
	"context"
	"notifier/domain/listener/models"
)

type Listener interface {
	AddQueueToListen(queueName string) error
	AddMessageHandler(typeMessage string, handler models.HandlerFunc)
	Listen(ctx context.Context)
	Stop()
}
