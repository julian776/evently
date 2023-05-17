package postgredb

import (
	"database/sql"
	"events-manager/pkgs/logger"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewPostgreRepository(l logger.Logger, settings *PostgreSettigs) *PostgreRepository {
	db, err := sql.Open("postgres", connToString(settings))
	if err != nil {
		l.Errorf("Error connecting to the DB: %s\n", err.Error())
	}

	// check if we can ping our DB
	err = db.Ping()
	if err != nil {
		l.Errorf("Error could not ping database: %s\n", err.Error())
	}

	return &PostgreRepository{
		db,
		l,
	}
}

// Take our connection struct and convert to a string for our db connection info
func connToString(info *PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)

}
