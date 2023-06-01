package events

import (
	"github.com/google/wire"
)

var UseCasesProvider = wire.NewSet(
	NewCreateEventUseCase,
	NewGetEventByIdUseCase,
	NewDeleteEventByIdUseCase,
	NewUpdateEventUseCase,
	NewGetAllEventsUseCase,
	NewAddAttendeeEventUseCase,
)
