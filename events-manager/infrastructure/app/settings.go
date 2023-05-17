package app

import "main/infrastructure/rabbit"

type AppSettings struct {
	Rabbit *rabbit.Settings
}
