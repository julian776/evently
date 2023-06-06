package repositories

import (
	"context"
	"events-manager/domain/users/models"
)

type UsersRepository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	CheckPasswordWithEmail(
		ctx context.Context,
		email string,
		passwordReceived string,
	) (bool, error)
}
