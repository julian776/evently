package postgredb

import (
	"context"
	"database/sql"
	"events-manager/domain/users/models"
	configs "events-manager/infrastructure/configs/postgres"
	"events-manager/pkgs/logger"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreUsersRepository struct {
	db *sql.DB
}

func NewPostgreUsersRepository(l logger.Logger, settings configs.PostgreSettigs) *PostgreUsersRepository {
	db, err := sql.Open("postgres", createConnToString(settings))
	if err != nil {
		l.Errorf("Error connecting [PostgreUsersRepository] to the DB: %s\n", err.Error())
	}

	// check if we can ping our DB
	err = db.Ping()
	if err != nil {
		l.Errorf("Error [PostgreUsersRepository] could not ping database: %s\n", err.Error())
	}

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
	query := `select
email,
password,
name,
purpouseOfUse FROM users WHERE email=$1`

	row, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return false, models.User{}, err
	}

	defer row.Close()

	if row.Next() {
		var name, email, purpouseOfUse, passwordDB string
		err := row.Scan(
			&email,
			&passwordDB,
			&name,
			&purpouseOfUse,
		)
		if err != nil {
			return false, models.User{}, err
		}

		if passwordReceived == passwordDB {
			user = models.User{
				Name:         name,
				Email:        email,
				PurposeOfUse: purpouseOfUse,
			}
			return true, user, nil
		}
	}

	return false, models.User{}, nil
}

func (r *PostgreUsersRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (models.User, error) {
	var user models.User

	query := `select
email,
name,
purpouseOfUse FROM users WHERE email=$1`

	row, err := r.db.QueryContext(ctx, query, email)
	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		var name, email, purpouseOfUse string
		err := row.Scan(
			&email,
			&name,
			&purpouseOfUse,
		)
		if err != nil {
			return models.User{}, err
		}

		user = models.User{
			Name:         name,
			Email:        email,
			PurposeOfUse: purpouseOfUse,
		}
	}

	return user, nil
}

func (r *PostgreUsersRepository) CreateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	query := `INSERT INTO
users(email, name, password, purpouseOfUse)
values($1, $2, $3, $4) RETURNING email, name, purpouseOfUse;`

	var name, email, purpouseOfUse string
	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Name,
		user.Password,
		user.PurposeOfUse,
	).Scan(
		&email,
		&name,
		&purpouseOfUse,
	)

	userCreated := models.User{
		Name:         name,
		Email:        email,
		PurposeOfUse: purpouseOfUse,
	}

	if err != nil {
		return models.User{}, err
	}

	return userCreated, nil
}

func (r *PostgreUsersRepository) UpdateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	query := `update users
set name=$1, 
email=$2,
password=$3,
purpouseOfUse=$4 where id=$8;`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
		user.PurposeOfUse,
	)
	if err != nil {
		return models.User{}, err
	}

	return models.User{}, nil
}

// Take our connection struct and convert to a string for our db connection info
func createConnToString(info configs.PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
