package rabbit

import (
	"context"
	"events-manager/infrastructure/rabbit/mappers"
	"events-manager/pkgs/logger"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitPublisher struct {
	logger logger.Logger
	Conn   *amqp.Connection
	Ch     *amqp.Channel
}

func (c *RabbitPublisher) Stop() {
	c.Conn.Close()
	c.Ch.Close()
}

// Returns a new `RabbitPublisher` instance
// with the connection and channel set up.
func NewRabbitPublisher(
	logger logger.Logger,
	settings Settings,
) *RabbitPublisher {
	if settings.Url == "" {
		logger.Errorf("Can not connect to RabbitMQ url is blank")
		return &RabbitPublisher{}
	}
	logger.Infof("Url: %s", settings.Url)

	conn, err := amqp.Dial(settings.Url)
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return &RabbitPublisher{
		logger,
		conn,
		ch,
	}
}

func (c *RabbitPublisher) QueueDeclare(queueName string) error {
	_, err := c.Ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %s", err.Error())
	}
	return nil
}

// Publishes a message to a specified queue in RabbitMQ.
// It maps the message to a byte array using the
// `mappers.MapStructToMessage` function and then publishes
// the message.
func (c *RabbitPublisher) PublishMessageWithContext(
	ctx context.Context,
	queueName string,
	message map[string]any,
	typeMessage string,
) error {
	bodyToPublish, err := mappers.MapStructToMessage(message, typeMessage)
	if err != nil {
		return fmt.Errorf("message can not be parsed")
	}
	err = c.Ch.PublishWithContext(
		ctx,
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		bodyToPublish,
	)
	return err
}
