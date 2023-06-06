// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"events-manager/domain/events/usecases"
	"events-manager/domain/users/usecases"
	"events-manager/infrastructure/app"
	"events-manager/infrastructure/events/adapters/postgre_db"
	http2 "events-manager/infrastructure/http/client"
	"events-manager/infrastructure/http/server"
	"events-manager/infrastructure/rabbit"
	postgredb2 "events-manager/infrastructure/users/adapters/postgre_db"
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
	postgreEventsRepository := postgredb.NewPostgreEventsRepository(sugaredLogger, postgreSettigs)
	eventsSettings := app.GetEventsSettings(appSettings)
	createEventUseCase := events.NewCreateEventUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository, eventsSettings)
	getEventByIdUseCase := events.NewGetEventByIdUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository)
	deleteEventByIdUseCase := events.NewDeleteEventByIdUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository, eventsSettings)
	updateEventUseCase := events.NewUpdateEventUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository, eventsSettings)
	getAllEventsUseCase := events.NewGetAllEventsUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository)
	addAttendeeEventUseCase := events.NewAddAttendeeEventUseCase(sugaredLogger, rabbitPublisher, postgreEventsRepository, eventsSettings)
	postgreUsersRepository := postgredb2.NewPostgreUsersRepository(sugaredLogger, postgreSettigs)
	usersSettings := app.GetUsersSettings(appSettings)
	createUserUseCase := users.NewCreateEventUseCase(sugaredLogger, rabbitPublisher, postgreUsersRepository, usersSettings)
	getUserByEmailUseCase := users.NewGetUserByEmailUseCase(sugaredLogger, rabbitPublisher, postgreUsersRepository, usersSettings)
	loginUserUseCase := users.NewLoginUserUseCase(sugaredLogger, rabbitPublisher, postgreUsersRepository, usersSettings)
	appApp := app.NewApp(sugaredLogger, engine, client, appSettings, rabbitPublisher, createEventUseCase, getEventByIdUseCase, deleteEventByIdUseCase, updateEventUseCase, getAllEventsUseCase, addAttendeeEventUseCase, createUserUseCase, getUserByEmailUseCase, loginUserUseCase)
	return appApp
}
