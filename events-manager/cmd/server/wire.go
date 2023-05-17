//go:build wireinject
// +build wireinject

package main

import (
	"events-manager/infrastructure/app"
	http "events-manager/infrastructure/http/server"
	"events-manager/infrastructure/rabbit"
	"events-manager/pkgs/logger"

	"github.com/google/wire"
	"go.uber.org/zap"
)

func CreateApp() *app.App {
	wire.Build(
		app.SettingsProvider,
		logger.NewLogger,
		wire.Bind(new(logger.Logger), new(*zap.SugaredLogger)),
		http.NewServer,
		rabbit.NewRabbitClient,
		app.NewApp,
	)
	return new(app.App)
}
