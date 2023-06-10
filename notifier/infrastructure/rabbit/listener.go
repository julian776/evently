package rabbit

import (
	"context"
	"fmt"
	"log"
	"notifier/domain/listener/models"
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
	queuesToListen []string

	// Key is the type of message
	handlers map[string][]models.HandlerFunc
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
		queuesToListen: []string{},
		handlers:       make(map[string][]models.HandlerFunc),
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

	l.queuesToListen = append(l.queuesToListen, queueName)
	return nil
}

// This method allows adding a new handler
// function for a specific message type.
func (l *RabbitListener) AddMessageHandler(
	typeMessage string,
	handler models.HandlerFunc,
) {
	l.handlers[typeMessage] = append(l.handlers[typeMessage], handler)
}

// Listens to the queues added and processes
// the messages received from them.
// It starts a goroutine for each queue to consume
// messages from it and then processes each message
// using the appropiate handlers registered for its type.
// The method blocks until the context is done.
func (l *RabbitListener) Listen(
	ctx context.Context,
) {
	l.logger.Infof("Listening for incoming messages...")
	for _, queue := range l.queuesToListen {
		go func(queue string) {
			cMessages, err := l.ch.Consume(queue, "", false, false, false, false, nil)
			if err != nil {
				l.logger.Errorf("error consuming queue: %s", queue)
			}

			for message := range cMessages {
				go func(message amqp.Delivery) {
					l.processMessage(ctx, message)
				}(message)
			}
		}(queue)
	}

	<-ctx.Done()
}

// Checks if there is a registered handler for
// the message type, and if not, it logs a warning
// and ignores the message.
// If there is a registered handlers, it maps the
// `amqp.Delivery` message to a domain `Message`
// using a mapper function, and then calls each
// registered handler.
func (l *RabbitListener) processMessage(ctx context.Context, message amqp.Delivery) {
	handlers, ok := l.handlers[message.Type]
	if !ok {
		l.logger.Warnf("ignoring message due to no handler registered, message type [%s]", message.Type)
		return
	}
	messageMapped, err := mappers.MapAmqpToMessage(message)
	if err != nil {
		l.logger.Errorf("can not process message %s", err.Error())
	}

	for _, handler := range handlers {
		err := handler(ctx, messageMapped)
		if err != nil {
			l.logger.Errorf("error in handler for type [%s] when processing message %s", message.Type, err.Error())
			message.Acknowledger.Nack(message.DeliveryTag, false, false)
		}
	}

	message.Acknowledger.Ack(message.DeliveryTag, false)
}
