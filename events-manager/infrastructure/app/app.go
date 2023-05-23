package app

import (
	"context"
	"events-manager/domain/broker"
	events "events-manager/domain/events/usecases"
	"events-manager/pkgs/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
	Logger          logger.Logger
	Server          *gin.Engine
	Client          *http.Client
	Settings        AppSettings
	BrokerPublisher broker.BrokerPublisher

	//------ UseCases-------
	CreateEventUseCase *events.CreateEventUseCase
}

func NewApp(
	logger logger.Logger,
	server *gin.Engine,
	client *http.Client,
	appSettings AppSettings,
	brokerPublisher broker.BrokerPublisher,
	createEventUseCase *events.CreateEventUseCase,
) *App {
	return &App{
		logger,
		server,
		client,
		appSettings,
		brokerPublisher,
		createEventUseCase,
	}
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.Server,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	a.Logger.Infof("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		a.Logger.Fatalf("Server Shutdown: %s", err.Error())
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")

	return nil
}
