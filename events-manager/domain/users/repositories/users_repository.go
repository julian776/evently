package events

import (
	"context"
	"events-manager/domain/events/models"
)

type UsersRepository interface {
	GetUser(ctx context.Context, id string) (models.Event, error)
	GetAllEventsById(ctx context.Context, id string) ([]models.Event, error)
	CreateUser(ctx context.Context, event models.Event) (models.Event, error)
	UpdateUser(ctx context.Context, event models.Event) (models.Event, error)
	DeleteUser(ctx context.Context, id string) (bool, error)
}
