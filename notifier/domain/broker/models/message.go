package models

import (
	"time"
)

type Message struct {
	// Unique identifier for the command
	Id string
	// Identifies the source of the command
	Source string
	// Identifies the target app
	Target string
	// The type of command
	Command string
	// Command generation time
	GenerationTime *time.Time
	// The command data
	Body map[string]any
}
