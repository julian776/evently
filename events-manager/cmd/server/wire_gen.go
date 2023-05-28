// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"events-manager/domain/events/usecases"
	"events-manager/infrastructure/app"
	"events-manager/infrastructure/events/adapters/postgre_db"
	http2 "events-manager/infrastructure/http/client"
	"events-manager/infrastructure/http/server"
	"events-manager/infrastructure/rabbit"
	"events-manager/pkgs/logger"
)

// Injectors from wire.go:

func CreateApp() *app.App {
	appSettings := app.LoadAppSettings()
	settings := app.GetLoggerSettings(appSettings)
	sugaredLogger := logger.NewLogger(settings)
	engine := http.NewServer()
	client := http2.NewClient()
	rabbitSettings := app.GetRabbitSettings(appSettings)
	rabbitPublisher := rabbit.NewRabbitPublisher(sugaredLogger, rabbitSettings)
	postgreSettigs := app.GetPostgreSettings(appSettings)
	postgreRepository := postgredb.NewPostgreRepository(sugaredLogger, postgreSettigs)
	eventsSettings := app.GetEventSettings(appSettings)
	createEventUseCase := events.NewCreateEventUseCase(sugaredLogger, rabbitPublisher, postgreRepository, eventsSettings)
	getEventByIdUseCase := events.NewGetEventByIdUseCase(sugaredLogger, rabbitPublisher, postgreRepository)
	deleteEventByIdUseCase := events.NewDeleteEventByIdUseCase(sugaredLogger, rabbitPublisher, postgreRepository, eventsSettings)
	updateEventUseCase := events.NewUpdateEventUseCase(sugaredLogger, rabbitPublisher, postgreRepository, eventsSettings)
	getAllEventsUseCase := events.NewGetAllEventsUseCase(sugaredLogger, rabbitPublisher, postgreRepository)
	appApp := app.NewApp(sugaredLogger, engine, client, appSettings, rabbitPublisher, createEventUseCase, getEventByIdUseCase, deleteEventByIdUseCase, updateEventUseCase, getAllEventsUseCase)
	return appApp
}
