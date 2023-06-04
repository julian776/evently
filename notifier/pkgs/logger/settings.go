package logger

type Settings struct {
	Enviroment string `envconfig:"ENVIROMENT" default:prod`
}
