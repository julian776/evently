package events

import (
	"context"
	"events-manager/domain/events/models"
)

type EventsRepository interface {
	Get(ctx context.Context, id string) (models.Event, error)
	GetAllByUserId(ctx context.Context, id string) ([]models.Event, error)
	Create(ctx context.Context, event models.Event) (models.Event, error)
	Update(ctx context.Context, event models.Event) (models.Event, error)
	Delete(ctx context.Context, id string) (bool, error)
}
