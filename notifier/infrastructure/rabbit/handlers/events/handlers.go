package events

import (
	"context"
	events "notifier/domain/events"
	eventsModels "notifier/domain/events/models"
	"notifier/domain/listener/models"
	"notifier/infrastructure/app"

	"github.com/mitchellh/mapstructure"
)

func SetHandlers(a *app.App) {
	eventsQueue := a.Settings.EventsSettings.Queue
	a.Listener.AddQueueToListen(eventsQueue)

	a.Listener.AddMessageHandler(
		events.EVENT_CREATED,
		func(ctx context.Context, message models.Message) {
			var event eventsModels.Event
			err := mapstructure.Decode(message.Body, &event)
			if err != nil {
				a.Logger.Errorf("Can not parse message to event")
				return
			}

			a.NotifyAndSaveReminderUseCase.Execute(ctx, event)
		},
	)
}
