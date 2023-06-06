package app

import (
	"fmt"
	"notifier/infrastructure/rabbit"
	"notifier/infrastructure/reminders"
	"notifier/pkgs/logger"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var SettingsProvider = wire.NewSet(
	LoadAppSettings,
	GetLoggerSettings,
	GetRabbitSettings,
	GetMongoSettings,
)

var appSettings *AppSettings

type AppSettings struct {
	Logger       *logger.Settings
	Rabbit       *rabbit.Settings
	MongoSettigs *reminders.MongoSettigs
}

func LoadAppSettings() AppSettings {
	if appSettings == nil {
		settings := AppSettings{}
		if err := envconfig.Process("", &settings); err != nil {
			fmt.Printf("Error loading envs: %s", err.Error())
		}

		appSettings = &settings
	}

	return *appSettings
}

func GetLoggerSettings(settings AppSettings) logger.Settings {
	return *settings.Logger
}

func GetRabbitSettings(settings AppSettings) rabbit.Settings {
	return *settings.Rabbit
}

func GetMongoSettings(settings AppSettings) reminders.MongoSettigs {
	return *settings.MongoSettigs
}

func loadConfigFile() (AppSettings, error) {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.AddConfigPath("./go_apps/hibot/conversationscounter/configs/server")
	vp.SetConfigName("configs")
	vp.SetConfigType("json")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Printf("Can not load config file %s", err)
	}
	var settings AppSettings
	err = vp.Unmarshal(&settings)
	if err != nil {
		return AppSettings{}, err
	}
	return settings, nil
}
