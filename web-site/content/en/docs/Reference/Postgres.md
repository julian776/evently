---
title: "Postgres"
linkTitle: "Postgres"
weight: 2
---

## Event
```GO
type Event struct {
	Id string

	Title string 

	Description string 

	Cost float32

	Location string 

	Attendees []string

	OrganizerName string 

	OrganizerEmail string 

	StartDate string 

	EndDate string 

	StartTime string

	EndTime string
}
```

## User
```GO
type User struct {
	Email string 

	Name string 

	Password string 

	PurposeOfUse string 
}
```

