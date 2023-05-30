package repositories

import (
	"context"
	"events-manager/domain/users/models"
)

type UsersRepository interface {
	GetUserById(ctx context.Context, id string) (models.User, error)
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
	DeleteUserById(ctx context.Context, id string) error
}
