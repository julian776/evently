package events

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/pkgs/logger"
)

type GetEventByIdUseCase struct {
	logger           logger.Logger
	publisher        broker.BrokerPublisher
	eventsRepository repositories.EventsRepository
}

// It retreives an event.
// If any error occurs during the process, it logs
// the error and returns an empty event and the error.
func (u *GetEventByIdUseCase) Execute(ctx context.Context, id string) (models.Event, error) {
	event, err := u.eventsRepository.GetEventById(ctx, id)
	if err != nil {
		u.logger.Errorf("Error fetching event %s", err.Error())
		return models.Event{}, err
	}

	return event, nil
}

func NewGetEventByIdUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	eventsRepository repositories.EventsRepository,
) *GetEventByIdUseCase {
	return &GetEventByIdUseCase{
		logger,
		publisher,
		eventsRepository,
	}
}
