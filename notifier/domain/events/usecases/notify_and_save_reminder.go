package events

import (
	"context"
	"fmt"
	"notifier/domain/events/models"
	"notifier/domain/listener"
	rModels "notifier/domain/reminders/models"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/logger"
	"time"
)

type NotifyAndSaveReminderUseCase struct {
	logger              logger.Logger
	listener            listener.Listener
	remindersRepository reminders.RemindersRepository
}

func (u *NotifyAndSaveReminderUseCase) Execute(
	ctx context.Context,
	event models.Event,
) error {
	dateToSend, err := time.Parse("02/01/2006", event.StartDate)
	if err != nil {
		return err
	}

	reminder := rModels.Reminder{
		EventId:        event.Id,
		DateToSend:     dateToSend.String(),
		MessageToSend:  fmt.Sprintf("Remember your event %s", event.Title),
		EmailsToNotify: []string{},
	}

	reminderCreated, err := u.remindersRepository.CreateReminder(ctx, reminder)
	if err != nil {
		return err
	}
	fmt.Println("HPTA:", reminderCreated)

	return nil
}

func NewNotifyAndSaveReminderUseCase(
	logger logger.Logger,
	listener listener.Listener,
	remindersRepository reminders.RemindersRepository,
) *NotifyAndSaveReminderUseCase {
	return &NotifyAndSaveReminderUseCase{
		logger,
		listener,
		remindersRepository,
	}
}
