package app

import (
	"context"
	"log"
	"notifier/domain/broker"
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
	BrokerListener      broker.BrokerListener
	RemindersRepository reminders.RemindersRepository

	//------ UseCases-------

	//------ End UseCases-------

}

func NewApp(
	logger logger.Logger,
	appSettings AppSettings,
	brokerListener broker.BrokerListener,
	remindersRepository reminders.RemindersRepository,
) *App {
	return &App{
		logger,
		appSettings,
		brokerListener,
		remindersRepository,
	}
}

func (a *App) Run() error {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-ctx.Done()
	log.Println("timeout of 5 seconds. closing broker connection")

	// Close broker connection
	a.BrokerListener.Stop()

	log.Println("Server exiting")

	return nil
}
