package app

import (
	"context"
	"events-manager/infrastructure/rabbit"
	"events-manager/pkgs/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	Logger logger.Logger
	Server *gin.Engine
	//Client   *http.Client
	Settings     AppSettings
	RabbitClient *rabbit.RabbitClient
}

func NewApp(
	logger logger.Logger,
	server *gin.Engine,
	//client *http.Client,
	appSettings AppSettings,
	rabbitClient *rabbit.RabbitClient,
) *App {
	return &App{
		logger,
		server,
		//client,
		appSettings,
		rabbitClient,
	}

}

func (a *App) Run(ctx context.Context) error {
	a.Server.Run(":8080")
	return nil
}
