package app

import (
	configs "events-manager/infrastructure/configs/postgres"
	"events-manager/infrastructure/events"
	"events-manager/infrastructure/rabbit"
	"events-manager/infrastructure/users"
	"events-manager/pkgs/logger"
	"fmt"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var SettingsProvider = wire.NewSet(
	LoadAppSettings,
	GetLoggerSettings,
	GetPostgreSettings,
	GetRabbitSettings,
	GetEventsSettings,
	GetUsersSettings,
)

var appSettings *AppSettings

type AppSettings struct {
	Logger         *logger.Settings
	Rabbit         *rabbit.Settings
	PostgreSettigs *configs.PostgreSettigs
	EventsSettings *events.EventsSettings
	UsersSettings  *users.UsersSettings
	Port           uint64 `envconfig:"PORT" required:"true"`
	ApiKey         string `envconfig:"API_KEY" default:""`
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

func GetPostgreSettings(settings AppSettings) configs.PostgreSettigs {
	return *settings.PostgreSettigs
}

func GetEventsSettings(settings AppSettings) events.EventsSettings {
	return *settings.EventsSettings
}

func GetUsersSettings(settings AppSettings) users.UsersSettings {
	return *settings.UsersSettings
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
