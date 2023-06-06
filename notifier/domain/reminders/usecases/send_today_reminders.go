package reminders

// import (
// 	"context"
// 	"notifier/domain/reminders"
// 	"notifier/pkgs/logger"
// )

// type SendTodayRemindersUseCase struct {
// 	logger              logger.Logger
// 	remindersRepository reminders.RemindersRepository

// }

// func (u *SendTodayRemindersUseCase) Execute(ctx context.Context) error {
// 	todayReminders, err := u.remindersRepository.GetAllTodayReminders(ctx)
// 	if err != nil {
// 		u.logger.Errorf("error creating event %s", err.Error())
// 		return err
// 	}

// 	err = u.publisher.PublishMessageWithContext(
// 		ctx,
// 		u.eventsSettings.Queue,
// 		eventCreated,
// 		models.EVENT_CREATED,
// 	)
// 	if err != nil {
// 		u.logger.Errorf("Error publishing event %s", err.Error())
// 		return models.Event{}, err
// 	}

// 	return eventCreated, nil
// }

// func NewSendTodayRemindersUseCase(
// 	logger logger.Logger,
// 	publisher broker.BrokerPublisher,
// 	eventsRepository repositories.EventsRepository,
// 	eventsSettings events.EventsSettings,
// ) *SendTodayRemindersUseCase {
// 	return &SendTodayRemindersUseCase{
// 		logger,
// 		publisher,
// 		eventsRepository,
// 		eventsSettings,
// 	}
// }
