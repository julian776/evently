package reminders

import (
	"context"
	"notifier/domain/reminders/models"
)

type RemindersRepository interface {
	CreateReminder(
		ctx context.Context,
		reminder models.Reminder,
	) ([]models.Reminder, error)

	GetAllTodayReminders(ctx context.Context) ([]models.Reminder, error)

	AddNewEmailToReminderWithEventId(
		ctx context.Context,
		eventId string,
		email string,
	) error

	DeletePastReminders(ctx context.Context) error
}
