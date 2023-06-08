package events

import "github.com/google/wire"

var UseCasesProvider = wire.NewSet(
	NewNotifyAndSaveReminderUseCase,
)
