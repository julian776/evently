package events

import (
	"context"
	"events-manager/domain/broker"
	brokerMocks "events-manager/domain/broker/mocks"
	"events-manager/domain/events/models"
	"events-manager/domain/events/repositories"
	"events-manager/domain/events/repositories/mocks"
	"events-manager/infrastructure/events"
	"events-manager/pkgs/logger"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestCreateEventUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	mockRepo, mockBroker := setUpMocks()

	log := logger.NewLogger(logger.Settings{})

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
				logger:           log,
				eventsRepository: mockRepo,
				eventsSettings:   events.EventsSettings{},
				publisher:        mockBroker,
			},
			args: args{
				context.Background(),
				models.Event{},
			},
			wantErr: false,
			want:    models.Event{},
		},
		{
			name: "should fail with error from repo",
			fields: fields{
				logger:           log,
				eventsRepository: mockRepo,
				eventsSettings:   events.EventsSettings{},
				publisher:        mockBroker,
			},
			args: args{
				context.Background(),
				models.Event{Cost: -2.0},
			},
			wantErr: true,
			want:    models.Event{},
		},
		{
			name: "should fail with error from publisher",
			fields: fields{
				logger:           log,
				eventsRepository: mockRepo,
				eventsSettings:   events.EventsSettings{},
				publisher:        mockBroker,
			},
			args: args{
				context.Background(),
				models.Event{Id: "2"},
			},
			wantErr: true,
			want:    models.Event{},
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
				assert.Error(err)
				return
			}
			assert.Equal(tt.want, got, "different events were provided")
		})
	}

	mockBroker.AssertNumberOfCalls(t, "PublishMessageWithContext", 2)
	mockRepo.AssertNumberOfCalls(t, "CreateEvent", 3)
}

func setUpMocks() (*mocks.EventsRepository, *brokerMocks.BrokerPublisher) {
	mockBroker := &brokerMocks.BrokerPublisher{}
	mockBroker.On(
		"PublishMessageWithContext",
		context.Background(),
		"",
		map[string]any{"event": models.Event{}},
		"notifier.events.created",
	).Return(nil)
	mockBroker.On(
		"PublishMessageWithContext",
		context.Background(),
		"",
		map[string]any{"event": models.Event{Id: "2"}},
		"notifier.events.created",
	).Return(fmt.Errorf("Error"))

	// ----- REPO ------
	mockRepo := &mocks.EventsRepository{}
	mockRepo.On(
		"CreateEvent",
		context.Background(),
		models.Event{},
	).Return(models.Event{}, nil)
	mockRepo.On(
		"CreateEvent",
		context.Background(),
		models.Event{Id: "2"},
	).Return(models.Event{Id: "2"}, nil)
	mockRepo.On(
		"CreateEvent",
		context.Background(),
		models.Event{Cost: -2.0},
	).Return(models.Event{}, fmt.Errorf("error in repo"))

	return mockRepo, mockBroker
}
