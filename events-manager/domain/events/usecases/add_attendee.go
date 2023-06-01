package events

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/events/dtos"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/infrastructure/events"
	"events-manager/pkgs/logger"
)

type AddAttendeeEventUseCase struct {
	logger           logger.Logger
	publisher        broker.BrokerPublisher
	eventsRepository repositories.EventsRepository
	eventsSettings   events.EventsSettings
}

// It creates the event and publishes an event.
// If any error occurs during the process, it logs
// the error and returns an empty event and the error.
func (u *AddAttendeeEventUseCase) Execute(ctx context.Context, addAttendeDTO dtos.AddAttendeDTO) ([]string, error) {
	attendees, err := u.eventsRepository.AddAttendeToEventById(ctx, addAttendeDTO.EventId, addAttendeDTO.AttendeeEmail)
	if err != nil {
		u.logger.Errorf("Error creating event %s", err.Error())
		return []string{}, err
	}

	err = u.publisher.PublishMessageWithContext(
		ctx,
		u.eventsSettings.Queue,
		attendees,
		models.EVENT_CREATED,
	)
	if err != nil {
		u.logger.Errorf("Error publishing event %s", err.Error())
		return []string{}, err
	}

	return attendees, nil
}

func NewAddAttendeeEventUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	eventsRepository repositories.EventsRepository,
	eventsSettings events.EventsSettings,
) *AddAttendeeEventUseCase {
	return &AddAttendeeEventUseCase{
		logger,
		publisher,
		eventsRepository,
		eventsSettings,
	}
}
