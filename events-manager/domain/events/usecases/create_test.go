package events

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/domain/events/repositories/mocks"
	"events-manager/infrastructure/events"
	"events-manager/pkgs/logger"
	"reflect"
	"testing"
)

func TestCreateEventUseCase_Execute(t *testing.T) {
	mockRepo := &mocks.EventsRepository{}
	mockRepo.On("CreateEvent", context.Background(), models.Event{}).Return(models.Event{})

	type fields struct {
		logger           logger.Logger
		publisher        broker.BrokerPublisher
		eventsRepository repositories.EventsRepository
		eventsSettings   events.EventsSettings
	}
	type args struct {
		ctx   context.Context
		event models.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.Event
		wantErr bool
	}{
		{
			name: "should pass",
			fields: fields{
				eventsRepository: mockRepo,
			},
			args: args{
				context.Background(),
				models.Event{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &CreateEventUseCase{
				logger:           tt.fields.logger,
				publisher:        tt.fields.publisher,
				eventsRepository: tt.fields.eventsRepository,
				eventsSettings:   tt.fields.eventsSettings,
			}
			got, err := u.Execute(tt.args.ctx, tt.args.event)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEventUseCase.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEventUseCase.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
