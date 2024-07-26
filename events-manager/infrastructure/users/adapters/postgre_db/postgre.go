package postgredb

import (
	"context"
	"events-manager/domain/users/models"
	configs "events-manager/infrastructure/configs/postgres"
	"events-manager/pkgs/logger"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreUsersRepository struct {
	db *gorm.DB
}

func NewPostgreUsersRepository(l logger.Logger, settings configs.PostgreSettigs) *PostgreUsersRepository {
	db, err := gorm.Open(postgres.Open(createConnToString(settings)), &gorm.Config{})
	if err != nil {
		l.Errorf("Error connecting [PostgreUsersRepository] to the DB: %s\n", err.Error())
	}

	db.AutoMigrate(&models.User{})

	return &PostgreUsersRepository{
		db,
	}
}

// It queries the database to retrieve a
// user with the given email and checks
// if the passwordReceived matches the
// password stored in the database for that user.
func (r *PostgreUsersRepository) GetUserAndCheckPasswordWithEmail(
	ctx context.Context,
	email string,
	passwordReceived string,
) (bool, models.User, error) {
	var user models.User
	r.db.Where("email = ?", email).First(&user)

	if passwordReceived == user.Password {
		user.Password = ""
		return true, user, nil
	}

	return false, models.User{}, nil
}

func (r *PostgreUsersRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {
	log.Printf("email: %v\n", email)
	var user models.User

	r.db.Where("email = ?", email).Select("email, name, purpose_of_use").First(&user)

	log.Printf("user: %v\n", user)
	return user, nil
}

func (r *PostgreUsersRepository) CreateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	r.db.Create(&user)
	user.Password = ""

	return user, nil
}

func (r *PostgreUsersRepository) UpdateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	r.db.Save(&user)

	return models.User{}, nil
}

// Take our connection struct and convert to a string for our db connection info
func createConnToString(info configs.PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
