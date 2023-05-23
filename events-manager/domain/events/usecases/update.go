package events

import (
	"context"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
)

type UpdateUseCase struct {
	eventsRepository repositories.EventsRepository
}

func (u *UpdateUseCase) Execute(ctx context.Context, event models.Event) (models.Event, error) {
	eventCreated, err := u.eventsRepository.CreateEvent(ctx, event)
	if err != nil {
		return models.Event{}, err
	}
	return eventCreated, nil
}
