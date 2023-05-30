package users

import (
	"context"
	"events-manager/domain/broker"
	"events-manager/domain/users/models"
	"events-manager/domain/users/repositories"
	"events-manager/infrastructure/users"
	"events-manager/pkgs/logger"
)

type CreateUserUseCase struct {
	logger          logger.Logger
	publisher       broker.BrokerPublisher
	usersRepository repositories.UsersRepository
	userSettings    users.UsersSettings
}

// It creates the user and publishes an event.
// If any error occurs during the process, it logs
// the error and returns an empty user and the error.
func (u *CreateUserUseCase) Execute(ctx context.Context, user models.User) (models.User, error) {
	eventCreated, err := u.usersRepository.CreateUser(ctx, user)
	if err != nil {
		u.logger.Errorf("Error creating user %s", err.Error())
		return models.User{}, err
	}

	err = u.publisher.PublishMessageWithContext(
		ctx,
		u.userSettings.Queue,
		eventCreated,
		models.USER_CREATED,
	)
	if err != nil {
		u.logger.Errorf("Error publishing event %s", err.Error())
		return models.User{}, err
	}

	return eventCreated, nil
}

func NewCreateEventUseCase(
	logger logger.Logger,
	publisher broker.BrokerPublisher,
	usersRepository repositories.UsersRepository,
	usersSettings users.UsersSettings,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		logger,
		publisher,
		usersRepository,
		usersSettings,
	}
}
