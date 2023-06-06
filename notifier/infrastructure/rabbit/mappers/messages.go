package mappers

import (
	"encoding/json"
	"notifier/domain/broker/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MapAmqpToMessage(message amqp.Delivery) (models.Message, error) {
	body := make(map[string]any)
	err := json.Unmarshal(message.Body, &body)
	if err != nil {
		return models.Message{}, err
	}

	return models.Message{
		Id:             message.MessageId,
		GenerationTime: message.Timestamp,
		Type:           message.Type,
		Body:           body,
		Source:         message.AppId,
	}, nil
}
