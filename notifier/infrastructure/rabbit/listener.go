package rabbit

import (
	"context"
	"fmt"
	"log"
	"notifier/domain/broker/models"
	"notifier/infrastructure/rabbit/mappers"
	"notifier/pkgs/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitListener struct {
	logger         logger.Logger
	conn           *amqp.Connection
	ch             *amqp.Channel
	QueuesToListen []string

	// Key is the type of message
	Handlers map[string]models.HandlerFunc
}

func (l *RabbitListener) Stop() {
	l.ch.Close()
	l.conn.Close()
}

func NewRabbitListener(
	logger logger.Logger,
	settings Settings,
) *RabbitListener {
	if settings.Url == "" {
		logger.Errorf("Can not connect to RabbitMQ url is blank")
		return &RabbitListener{}
	}
	logger.Infof("Url: %s", settings.Url)

	conn, err := amqp.Dial(settings.Url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RabbitListener{
		logger:         logger,
		conn:           conn,
		ch:             ch,
		QueuesToListen: []string{},
	}
}

func (l *RabbitListener) AddQueueToListen(queueName string) error {
	_, err := l.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %s", err.Error())
	}

	l.QueuesToListen = append(l.QueuesToListen, queueName)
	return nil
}

func (l *RabbitListener) AddMessageHandler(
	typeMessage string,
	handler models.HandlerFunc,
) {
	l.Handlers[typeMessage] = handler
}

// Listens to the queues added to the `QueuesToListen`
// slice and processes the messages received from them.
// It starts a goroutine for each queue to consume
// messages from it and then processes each message
// using the appropriate handler registered for its type.
// The method blocks until the context is done.
func (l *RabbitListener) Listen(
	ctx context.Context,
) {
	for _, queue := range l.QueuesToListen {
		go func(queue string) {
			cMessages, err := l.ch.Consume(queue, "", false, false, false, false, nil)
			if err != nil {
				l.logger.Errorf("error consuming queue: %s", queue)
			}

			for message := range cMessages {
				l.processMessage(ctx, message)
			}
		}(queue)
	}

	<-ctx.Done()
}

func (l *RabbitListener) processMessage(ctx context.Context, message amqp.Delivery) {
	handler, ok := l.Handlers[message.Type]
	if !ok {
		l.logger.Warnf("ignoring message due to no handler registered, message type %s", message.Type)
	}
	messageMapped, err := mappers.MapAmqpToMessage(message)
	if err != nil {
		l.logger.Errorf("can not process message %s", err.Error())
	}
	handler(ctx, messageMapped)
}
