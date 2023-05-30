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
	fmt.Println(settings)
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

func (r *PostgreUsersRepository) GetUserById(
	ctx context.Context,
	id string,
) (models.User, error) {
	var user models.User

	query := `select
name,
email,
password,
purpouseOfUse FROM users WHERE id=$1`

	row, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		var name, email, password, purpouseOfUse string
		err := row.Scan(
			&name,
			&email,
			&password,
			&purpouseOfUse,
		)
		if err != nil {
			return models.User{}, err
		}

		user = models.User{
			Name:          name,
			Email:         email,
			Password:      password,
			PurpouseOfUse: purpouseOfUse,
		}
	}

	return user, nil
}

func (r *PostgreUsersRepository) CreateUser(
	ctx context.Context,
	user models.User,
) (models.User, error) {
	query := `insert into
users(email, name, password, purpouseOfUse)
values($1, $2, $3, $4) RETURNING *;`

	var name, email, password, purpouseOfUse string
	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Name,
		user.Password,
		user.PurpouseOfUse,
	).Scan(
		&email,
		&name,
		&password,
		&purpouseOfUse,
	)

	userCreated := models.User{
		Name:          name,
		Email:         email,
		Password:      password,
		PurpouseOfUse: purpouseOfUse,
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
		user.PurpouseOfUse,
	)
	if err != nil {
		return models.User{}, err
	}

	return models.User{}, nil
}

func (r *PostgreUsersRepository) DeleteUserById(
	ctx context.Context,
	id string,
) error {
	query := `delete from users where id=$1;`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

// Take our connection struct and convert to a string for our db connection info
func createConnToString(info configs.PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
