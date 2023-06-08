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
) {
	dateToSend, err := time.Parse("ISO", event.StartDate)
	fmt.Println("DATE: ", dateToSend)
	if err != nil {
		return
	}

	reminder := rModels.Reminder{
		EventId:        event.Id,
		DateToSend:     dateToSend.String(),
		MessageToSend:  fmt.Sprintf("Remember your event %s", event.Title),
		EmailsToNotify: []string{},
	}
	fmt.Println(reminder)
	//todayReminders, err := u.remindersRepository.CreateReminder(ctx)
	if err != nil {
		u.logger.Errorf("error creating event %s", err.Error())
		return
	}

	if err != nil {
		u.logger.Errorf("Error publishing event %s", err.Error())
		return
	}

	return
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
