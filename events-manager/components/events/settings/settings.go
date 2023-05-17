package configs

import (
	"fmt"
	"main/pkgs/logger"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type EventSettings struct {
	Logger  *logger.Settings
	Postgre *settings.MongoSettings
	Port    uint64 `envconfig:"PORT" required:"true"`
	ApiKey  string `envconfig:"API_KEY" default:""`
}

func LoadAppSettings() AppSettings {
	dir, _ := os.Getwd()
	fmt.Printf("Dir %s\n", dir)
	if appSettings == nil {
		settings := AppSettings{}
		if err := envconfig.Process("", &settings); err != nil {
			fmt.Printf("Error loading envs: %s", err.Error())
		}

		appSettings = &settings
	}

	return *appSettings
}
