package users

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/users/models"
	"events-manager/domain/users/repositories"
	"events-manager/infrastructure/users"
	"events-manager/pkgs/logger"
)

type GetUserByEmailUseCase struct {
	logger          logger.Logger
	publisher       broker.BrokerPublisher
	usersRepository repositories.UsersRepository
	userSettings    users.UsersSettings
}

// It creates the user and publishes an event.
// If any error occurs during the process, it logs
// the error and returns an empty user and the error.
func (u *GetUserByEmailUseCase) Execute(ctx context.Context, email string) (models.User, error) {
	user, err := u.usersRepository.GetUserByEmail(ctx, email)
	if err != nil {
		u.logger.Errorf("Error fetching user %s", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func NewGetUserByEmailUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	usersRepository repositories.UsersRepository,
	usersSettings users.UsersSettings,
) *GetUserByEmailUseCase {
	return &GetUserByEmailUseCase{
		logger,
		publisher,
		usersRepository,
		usersSettings,
	}
}
