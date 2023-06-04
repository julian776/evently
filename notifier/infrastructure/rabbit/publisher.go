package rabbit

import (
	"context"
	"events-manager/domain/broker/models"
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

func (c *RabbitPublisher) PublishMessageWithContext(
	ctx context.Context,
	queueName string,
	message any,
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

// Never used (Practicing) TODO: Delete
func (c *RabbitPublisher) ListenMessages(
	ctx context.Context,
	queueName string,
) chan<- models.Message {
	cMessages := make(chan models.Message)
	go func() {
		for message := range cMessages {
			err := c.PublishMessageWithContext(ctx, queueName, message, "")
			if err != nil {
				c.logger.Errorf("error sending message %s", message)
			}
		}
	}()
	return cMessages
}
