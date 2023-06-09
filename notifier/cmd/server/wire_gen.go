// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"notifier/domain/events/usecases"
	"notifier/infrastructure/app"
	"notifier/infrastructure/rabbit"
	"notifier/infrastructure/reminders"
	"notifier/pkgs/logger"
)

// Injectors from wire.go:

func CreateApp() *app.App {
	appSettings := app.LoadAppSettings()
	settings := app.GetLoggerSettings(appSettings)
	sugaredLogger := logger.NewLogger(settings)
	rabbitSettings := app.GetRabbitSettings(appSettings)
	rabbitListener := rabbit.NewRabbitListener(sugaredLogger, rabbitSettings)
	mongoSettigs := app.GetMongoSettings(appSettings)
	remindersMongoRepository := reminders.NewRemindersMongoRepository(sugaredLogger, mongoSettigs)
	emailsSettings := app.GetEmailsSettings(appSettings)
	notifyAndSaveReminderUseCase := events.NewNotifyAndSaveReminderUseCase(sugaredLogger, remindersMongoRepository, emailsSettings)
	notifyNewAttendeeAndUpdateReminderUseCase := events.NewNotifyNewAttendeeAndUpdateReminderUseCase(sugaredLogger, rabbitListener, remindersMongoRepository, emailsSettings)
	appApp := app.NewApp(sugaredLogger, appSettings, rabbitListener, remindersMongoRepository, notifyAndSaveReminderUseCase, notifyNewAttendeeAndUpdateReminderUseCase)
	return appApp
}
