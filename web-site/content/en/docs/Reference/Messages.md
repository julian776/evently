---
title: "Messages"
linkTitle: "Messages"
weight: 2
---

## Message
```GO
type Message struct {
	// Unique identifier for the command
	Id string
	// Identifies the source of the command
	Source string
	// Identifies the target app
	Target string
	// The type of command
	Type string
	// Command generation time
	GenerationTime time.Time
	// The command data
	Body map[string]any
}
```

