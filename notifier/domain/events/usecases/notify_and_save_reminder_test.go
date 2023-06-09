package events

import (
	"context"
	"notifier/domain/events/models"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/emails"
	"notifier/pkgs/logger"
	"testing"
)

func TestNotifyAndSaveReminderUseCase_Execute(t *testing.T) {
	type fields struct {
		logger              logger.Logger
		remindersRepository reminders.RemindersRepository
		emailsSettings      emails.Settings
	}
	type args struct {
		ctx   context.Context
		event models.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &NotifyAndSaveReminderUseCase{
				logger:              tt.fields.logger,
				remindersRepository: tt.fields.remindersRepository,
				emailsSettings:      tt.fields.emailsSettings,
			}
			if err := u.Execute(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("NotifyAndSaveReminderUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
