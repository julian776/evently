package users

import "github.com/google/wire"

var UseCasesProvider = wire.NewSet(
	NewCreateEventUseCase,
	NewGetUserByEmailUseCase,
)
