package repositories

import (
	"context"
	"events-manager/domain/events/models"
)

type EventsRepository interface {
	GetEventById(ctx context.Context, id string) (models.Event, error)
	GetAllEventsByOrganizerEmail(ctx context.Context, id string) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	UpdateEvent(ctx context.Context, event models.Event) (models.Event, error)
	DeleteEventById(ctx context.Context, id string) (models.Event, error)
}
