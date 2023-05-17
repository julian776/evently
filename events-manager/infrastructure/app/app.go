package app

import (
	"main/pkgs/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Logger *logger.Logger
	Server *gin.Engine
	Client *http.Client
	// Settings *settings.AppSettings
}

func NewApp(
	logger *logger.Logger,
	server *gin.Engine,
	client *http.Client,
) *App {
	return &App{
		logger,
		server,
		client,
	}

}
