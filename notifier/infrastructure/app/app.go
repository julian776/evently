package app

import (
	"context"
	"log"
	events "notifier/domain/events/usecases"
	"notifier/domain/listener"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/logger"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	Logger              logger.Logger
	Settings            AppSettings
	Listener            listener.Listener
	RemindersRepository reminders.RemindersRepository

	//------ UseCases-------
	NotifyAndSaveReminderUseCase              *events.NotifyAndSaveReminderUseCase
	NotifyNewAttendeeAndUpdateReminderUseCase *events.NotifyNewAttendeeAndUpdateReminderUseCase
	//------ End UseCases-------

}

func NewApp(
	logger logger.Logger,
	appSettings AppSettings,
	listener listener.Listener,
	remindersRepository reminders.RemindersRepository,
	notifyAndSaveReminderUseCase *events.NotifyAndSaveReminderUseCase,
	notifyNewAttendeeAndUpdateReminderUseCase *events.NotifyNewAttendeeAndUpdateReminderUseCase,
) *App {
	return &App{
		logger,
		appSettings,
		listener,
		remindersRepository,
		notifyAndSaveReminderUseCase,
		notifyNewAttendeeAndUpdateReminderUseCase,
	}
}

func (a *App) Run() error {

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	a.Listener.Listen(ctx)
	a.Logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-ctx.Done()
	log.Println("timeout of 5 seconds. closing broker connection")

	// Close broker connection
	a.Listener.Stop()

	log.Println("Server exiting")

	return nil
}
