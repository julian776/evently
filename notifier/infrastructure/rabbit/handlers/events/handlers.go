package events

import (
	"context"
	"fmt"
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
		func(ctx context.Context, message models.Message) error {
			var event eventsModels.Event
			err := mapstructure.Decode(message.Body["event"], &event)
			if err != nil {
				return fmt.Errorf("can not parse message to event %s", err.Error())
			}

			return a.NotifyAndSaveReminderUseCase.Execute(ctx, event)
		},
	)

	a.Listener.AddMessageHandler(
		events.ADDED_ATTENDEE,
		func(ctx context.Context, message models.Message) error {
			return a.NotifyNewAttendeeAndUpdateReminderUseCase.Execute(
				ctx,
				message.Body["eventId"].(string),
				message.Body["attendee"].(string),
			)
		},
	)
}
