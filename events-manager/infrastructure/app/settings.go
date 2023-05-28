package app

import (
	"events-manager/infrastructure/events"
	postgredb "events-manager/infrastructure/events/adapters/postgre_db"
	"events-manager/infrastructure/rabbit"
	"events-manager/pkgs/logger"
	"fmt"
	"os"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var SettingsProvider = wire.NewSet(
	LoadAppSettings,
	GetLoggerSettings,
	GetPostgreSettings,
	GetRabbitSettings,
	GetEventSettings,
)

var appSettings *AppSettings

type AppSettings struct {
	Logger         *logger.Settings
	Rabbit         *rabbit.Settings
	PostgreSettigs *postgredb.PostgreSettigs
	EventsSettings *events.EventsSettings
	Port           uint64 `envconfig:"PORT" required:"true"`
	ApiKey         string `envconfig:"API_KEY" default:""`
}

func LoadAppSettings() AppSettings {
	dir, _ := os.Getwd()
	fmt.Printf("Dir %s\n", dir)

	fmt.Println("Env: ", os.Getenv("RABBIT_URL"))
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

func GetPostgreSettings(settings AppSettings) postgredb.PostgreSettigs {
	return *settings.PostgreSettigs
}

func GetEventSettings(settings AppSettings) events.EventsSettings {
	return *settings.EventsSettings
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