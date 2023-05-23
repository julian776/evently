package app

import (
	"context"
	"events-manager/domain/broker"
	events "events-manager/domain/events/usecases"
	"events-manager/pkgs/logger"
	"net/http"

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

func (a *App) Run(ctx context.Context) error {
	a.Server.Run(":8080")
	return nil
}
