package app

import (
	"context"
	"log"
	events "notifier/domain/events/usecases"
	"notifier/domain/listener"
	reminders "notifier/domain/reminders/repositories"
	"notifier/pkgs/logger"
	"os"
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
	NotifyAndSaveReminderUseCase *events.NotifyAndSaveReminderUseCase
	//------ End UseCases-------

}

func NewApp(
	logger logger.Logger,
	appSettings AppSettings,
	listener listener.Listener,
	remindersRepository reminders.RemindersRepository,
	notifyAndSaveReminderUseCase *events.NotifyAndSaveReminderUseCase,
) *App {
	return &App{
		logger,
		appSettings,
		listener,
		remindersRepository,
		notifyAndSaveReminderUseCase,
	}
}

func (a *App) Run() error {
	a.Listener.Listen(context.TODO())

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
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
