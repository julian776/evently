package repositories

import (
	"context"
	"events-manager/domain/events/models"
)

//go:generate mockery --name=EventsRepository
type EventsRepository interface {
	GetEventById(ctx context.Context, id string) (models.Event, error)
	GetAllEvents(ctx context.Context) ([]models.Event, error)
	GetAllEventsByOrganizerEmail(ctx context.Context, id string) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	AddAttendeToEventById(
		ctx context.Context,
		id string,
		attendeeEmail string,
	) ([]string, error)
	UpdateEvent(ctx context.Context, event models.Event) (models.Event, error)
	DeleteEventById(ctx context.Context, id string) (models.Event, error)
}
