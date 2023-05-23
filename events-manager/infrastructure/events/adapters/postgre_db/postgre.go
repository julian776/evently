package postgredb

import (
	"context"
	"database/sql"
	"events-manager/domain/events/models"
	"events-manager/pkgs/logger"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgreRepository struct {
	db *sql.DB
}

func NewPostgreRepository(l logger.Logger, settings PostgreSettigs) *PostgreRepository {
	fmt.Println(settings)
	db, err := sql.Open("postgres", createConnToString(settings))
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
	}
}

func (r *PostgreRepository) GetEventById(
	ctx context.Context,
	id string,
) (models.Event, error) {
	var article models.Event

	query := `select title, content from events where id=$1`
	row, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return article, err
	}

	defer row.Close()

	if row.Next() {
		var title, content string

		err := row.Scan(&title, &content)
		if err != nil {
			return article, err
		}

		article = models.Event{
			Id:          id,
			Title:       title,
			Description: content,
		}
	}

	return models.Event{}, nil
}

func (r *PostgreRepository) CreateEvent(
	ctx context.Context,
	event models.Event,
) (models.Event, error) {
	query := `insert into events(title, description, location, organizerName, organizerEmail, startTime, endTime) values($1, $2, $3, $4, $5, $6, $7);`

	_, err := r.db.ExecContext(
		ctx,
		query,
		event.Title,
		event.Description,
		event.Location,
		event.OrganizerName,
		event.OrganizerEmail,
		event.StartTime,
		event.EndTime,
	)

	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (r *PostgreRepository) GetAllEventsByOrganizerEmail(
	ctx context.Context,
	id string,
) ([]models.Event, error) {
	return []models.Event{}, nil
}

func (r *PostgreRepository) UpdateEvent(
	ctx context.Context,
	event models.Event,
) (models.Event, error) {
	query := `update events set title=$1, content=$2 where id=$3;`

	_, err := r.db.ExecContext(ctx, query, event.Title, event.Description, event.Id)
	if err != nil {
		return models.Event{}, err
	}

	return models.Event{}, nil
}

func (r *PostgreRepository) DeleteEventById(
	ctx context.Context,
	id string,
) (models.Event, error) {
	query := `delete from events where id=$1;`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return models.Event{}, err
	}
	return models.Event{}, nil
}

// Take our connection struct and convert to a string for our db connection info
func createConnToString(info PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
