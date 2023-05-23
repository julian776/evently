package events

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/infrastructure/events"
	"events-manager/pkgs/logger"
)

type DeleteEventByIdUseCase struct {
	logger           logger.Logger
	publisher        broker.BrokerPublisher
	eventsRepository repositories.EventsRepository
	eventsSettings   events.EventsSettings
}

// It deletes the event and publishes the event.
// If any error occurs during the process, it logs
// the error and returns an empty event and the error.
func (u *DeleteEventByIdUseCase) Execute(ctx context.Context, id string) (models.Event, error) {
	eventCreated, err := u.eventsRepository.DeleteEventById(ctx, id)
	if err != nil {
		u.logger.Errorf("error deleting event %s", err.Error())
		return models.Event{}, err
	}

	err = u.publisher.PublishMessageWithContext(
		ctx,
		u.eventsSettings.Queue,
		eventCreated,
		models.EVENT_DELETED,
	)
	if err != nil {
		u.logger.Errorf("Error publishing delete event %s", err.Error())
		return models.Event{}, err
	}

	return eventCreated, nil
}

func NewDeleteEventByIdUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	eventsRepository repositories.EventsRepository,
	eventsSettings events.EventsSettings,
) *DeleteEventByIdUseCase {
	return &DeleteEventByIdUseCase{
		logger,
		publisher,
		eventsRepository,
		eventsSettings,
	}
}
