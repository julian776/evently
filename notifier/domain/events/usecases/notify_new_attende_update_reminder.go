package events

import (
	"context"
	"fmt"
	"notifier/domain/listener"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/emails"
	"notifier/pkgs/logger"
)

type NotifyNewAttendeeAndUpdateReminderUseCase struct {
	logger              logger.Logger
	listener            listener.Listener
	remindersRepository reminders.RemindersRepository
	emailsSettings      emails.Settings
}

func (u *NotifyNewAttendeeAndUpdateReminderUseCase) Execute(
	ctx context.Context,
	eventId string,
	attedeeEmail string,
) error {
	auth := emails.Auth{
		Email:    u.emailsSettings.Email,
		Password: u.emailsSettings.Password,
	}
	message := []byte(fmt.Sprintf("Evently\n\nYou registered succesfully to the event\n\n%s"))

	err := emails.SendEmail(
		ctx,
		auth,
		[]string{attedeeEmail},
		message,
	)
	if err != nil {
		return err
	}

	err = u.remindersRepository.AddNewEmailToReminderWithEventId(ctx, eventId, attedeeEmail)
	if err != nil {
		return err
	}

	return nil
}

func NewNotifyNewAttendeeAndUpdateReminderUseCase(
	logger logger.Logger,
	listener listener.Listener,
	remindersRepository reminders.RemindersRepository,
	emailsSettings emails.Settings,
) *NotifyNewAttendeeAndUpdateReminderUseCase {
	return &NotifyNewAttendeeAndUpdateReminderUseCase{
		logger,
		listener,
		remindersRepository,
		emailsSettings,
	}
}
