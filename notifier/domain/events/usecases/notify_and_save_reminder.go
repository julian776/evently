package events

import (
	"context"
	"fmt"
	"notifier/domain/events/models"
	rModels "notifier/domain/reminders/models"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/emails"
	"notifier/pkgs/logger"
	"time"
)

type NotifyAndSaveReminderUseCase struct {
	logger              logger.Logger
	remindersRepository reminders.RemindersRepository
	emailsSettings      emails.Settings
}

func (u *NotifyAndSaveReminderUseCase) Execute(
	ctx context.Context,
	event models.Event,
) error {
	auth := emails.Auth{
		Email:    u.emailsSettings.Email,
		Password: u.emailsSettings.Password,
	}
	message := []byte(fmt.Sprintf("Evently\n\nYour new event was succesfully created\n\n%s", event.Title))

	err := emails.SendEmail(
		ctx,
		auth,
		[]string{event.OrganizerEmail},
		message,
	)
	if err != nil {
		return err
	}

	dateToSend, err := time.Parse("02/01/2006", event.StartDate)
	if err != nil {
		return err
	}

	reminder := rModels.Reminder{
		EventId:        event.Id,
		DateToSend:     dateToSend,
		MessageToSend:  fmt.Sprintf("Remember your event %s", event.Title),
		EmailsToNotify: []string{},
	}

	_, err = u.remindersRepository.CreateReminder(ctx, reminder)
	if err != nil {
		return err
	}

	return nil
}

func NewNotifyAndSaveReminderUseCase(
	logger logger.Logger,
	remindersRepository reminders.RemindersRepository,
	emailsSettings emails.Settings,
) *NotifyAndSaveReminderUseCase {
	return &NotifyAndSaveReminderUseCase{
		logger,
		remindersRepository,
		emailsSettings,
	}
}
