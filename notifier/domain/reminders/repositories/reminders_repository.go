package reminders

import (
	"context"
	"notifier/domain/reminders/models"
)

type RemindersRepository interface {
	GetAllTodayReminders(ctx context.Context) ([]models.Reminder, error)
	AddNewEmailToReminderWithEventId(ctx context.Context, eventId string, email string) error
	// GetAllEventsByOrganizerEmail(ctx context.Context, id string) ([]models.Event, error)
	// CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	// AddAttendeToEventById(
	// 	ctx context.Context,
	// 	id string,
	// 	attendeeEmail string,
	// ) ([]string, error)
	// UpdateEvent(ctx context.Context, event models.Event) (models.Event, error)
	DeletePastReminders(ctx context.Context) error
}
