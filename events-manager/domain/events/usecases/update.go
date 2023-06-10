package events

import (
	"context"
	"events-manager/domain/broker"
	types "events-manager/domain/events"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/infrastructure/events"
	"events-manager/pkgs/logger"
)

type UpdateEventUseCase struct {
	logger           logger.Logger
	publisher        broker.BrokerPublisher
	eventsRepository repositories.EventsRepository
	eventsSettings   events.EventsSettings
}

// It creates the event and publishes an event.
// If any error occurs during the process, it logs
// the error and returns an empty event and the error.
func (u *UpdateEventUseCase) Execute(ctx context.Context, event models.Event) (models.Event, error) {
	eventUpdated, err := u.eventsRepository.UpdateEvent(ctx, event)
	if err != nil {
		u.logger.Errorf("Error updating event %s", err.Error())
		return models.Event{}, err
	}

	err = u.publisher.PublishMessageWithContext(
		ctx, u.eventsSettings.Queue,
		map[string]any{"event": eventUpdated},
		types.EVENT_UPDATED,
	)
	if err != nil {
		u.logger.Errorf("Error publishing event %s", err.Error())
		return models.Event{}, err
	}

	return eventUpdated, nil
}

func NewUpdateEventUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	eventsRepository repositories.EventsRepository,
	eventsSettings events.EventsSettings,
) *UpdateEventUseCase {
	return &UpdateEventUseCase{
		logger,
		publisher,
		eventsRepository,
		eventsSettings,
	}
}
