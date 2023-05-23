package services

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/broker/models"
	events "events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/pkgs/logger"
)

type EventsServices struct {
	logger           logger.Logger
	publisher        *broker.BrokerPublisher
	cMessage         chan models.Message
	eventsRepository repositories.EventsRepository
}

func (s *EventsServices) GetEvent(ctx context.Context, id string) (events.Event, error) {
	eventCreated, err := s.eventsRepository.GetEventById(ctx, id)
	if err != nil {
		s.logger.Errorf("error getting event: %s", err)
	}
	return eventCreated, nil
}

func (s *EventsServices) CreateEvent(ctx context.Context, event events.Event) (events.Event, error) {
	eventCreated, err := s.eventsRepository.CreateEvent(ctx, event)
	if err != nil {
		s.logger.Errorf("error creating event: %s", err)
	}
	return eventCreated, nil
}

func (s *EventsServices) UpdateEvent(event events.Event) error {
	return nil
}
