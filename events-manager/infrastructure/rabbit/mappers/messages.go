package mappers

import (
	"encoding/json"
	domain "events-manager/domain/app"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func MapStructToMessage(structPayload map[string]any, typeMessage string) (amqp.Publishing, error) {
	bodyBytes, err := json.Marshal(structPayload)
	if err != nil {
		return amqp.Publishing{}, err
	}

	return amqp.Publishing{
		Headers: map[string]interface{}{
			"sourceApplication": domain.APP_NAME,
		},
		ContentType:     "application/json",
		ContentEncoding: "UTF-8",
		MessageId:       uuid.New().String(),
		Timestamp:       time.Now(),
		Type:            typeMessage,
		AppId:           domain.APP_NAME,
		Body:            bodyBytes,
	}, nil
}
