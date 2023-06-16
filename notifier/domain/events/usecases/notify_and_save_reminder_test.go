package events

import (
	"context"
	eModels "notifier/domain/events/models"
	"notifier/domain/reminders/models"
	reminders "notifier/domain/reminders/repositories"
	"notifier/domain/reminders/repositories/mocks"
	"notifier/pkgs/emails"
	"notifier/pkgs/logger"
	"testing"
	"time"
)

type fields struct {
	logger              logger.Logger
	remindersRepository reminders.RemindersRepository
	emailsSettings      emails.Settings
}
type args struct {
	ctx   context.Context
	event eModels.Event
}

func TestNotifyAndSaveReminderUseCase_Execute(t *testing.T) {

	mockRepo := setUpMocks()

	log := logger.NewLogger(logger.Settings{})

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should fail with no date",
			fields: fields{
				logger:              log,
				remindersRepository: mockRepo,
				emailsSettings:      emails.Settings{},
			},
			args: args{
				ctx:   context.Background(),
				event: eModels.Event{},
			},
			wantErr: true,
		},
		{
			name: "should fail with wrong email credentials",
			fields: fields{
				logger:              log,
				remindersRepository: mockRepo,
				emailsSettings:      emails.Settings{},
			},
			args: args{
				ctx:   context.Background(),
				event: eModels.Event{StartDate: "06/05/2023"},
			},
			wantErr: true,
		},
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

func setUpMocks() reminders.RemindersRepository {
	mockRepo := &mocks.RemindersRepository{}
	mockRepo.On(
		"CreateReminder",
		context.Background(),
		models.Reminder{},
	).Return(models.Reminder{}, nil)

	dateToSend := time.Date(2023, time.May, 6, 0, 0, 0, 0, time.UTC)
	mockRepo.On(
		"CreateReminder",
		context.Background(),
		models.Reminder{
			EventId:        "",
			DateToSend:     dateToSend,
			MessageToSend:  "Remember your event ",
			EmailsToNotify: []string{},
		},
	).Return(models.Reminder{}, nil)

	return mockRepo
}
