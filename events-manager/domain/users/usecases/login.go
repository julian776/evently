package users

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/users/dtos"
	"events-manager/domain/users/repositories"
	"events-manager/infrastructure/users"
	"events-manager/pkgs/logger"
)

type LoginUserUseCase struct {
	logger          logger.Logger
	publisher       broker.BrokerPublisher
	usersRepository repositories.UsersRepository
	userSettings    users.UsersSettings
}

// It creates the user and publishes an event.
// If any error occurs during the process, it logs
// the error and returns an empty user and the error.
func (u *LoginUserUseCase) Execute(ctx context.Context, login dtos.LoginDTO) (bool, error) {
	isValid, err := u.usersRepository.CheckPasswordWithEmail(ctx, login.Email, login.Password)
	if err != nil {
		u.logger.Errorf("Error validating password %s", err.Error())
		return false, err
	}

	return isValid, nil
}

func NewLoginUserUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	usersRepository repositories.UsersRepository,
	usersSettings users.UsersSettings,
) *LoginUserUseCase {
	return &LoginUserUseCase{
		logger,
		publisher,
		usersRepository,
		usersSettings,
	}
}
