package main

import (
	app "main/infrastructure/app"
	http "main/infrastructure/http/server"
	"main/pkgs/logger"

	"github.com/google/wire"
)

func CreateApp() *app.App {
	wire.Build(
		settings.SettingsProvider,
		logger.NewLogger,
		http.NewServer,
		app.NewApp,
	)
	return new(app.App)
}
