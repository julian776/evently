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
	var event models.Event

	query := `select title, description, location, organizerName, organizerEmail, startTime, endTime from events where id=$1`
	row, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return event, err
	}

	defer row.Close()

	if row.Next() {
		var title, description, location, organizerName, organizerEmail, startTime, endTime string

		err := row.Scan(
			&title,
			&description,
			&location,
			&organizerName,
			&organizerEmail,
			&startTime,
			&endTime,
		)
		if err != nil {
			return models.Event{}, err
		}

		event = models.Event{
			Id:             id,
			Title:          title,
			Description:    description,
			Location:       location,
			OrganizerName:  organizerName,
			OrganizerEmail: organizerEmail,
			StartTime:      startTime,
			EndTime:        endTime,
		}
	}

	return event, nil
}

func (r *PostgreRepository) GetAllEvents(
	ctx context.Context,
) ([]models.Event, error) {
	var events []models.Event

	query := `select id, title, description, location, organizerName, organizerEmail, startTime, endTime from events`
	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return []models.Event{}, err
	}

	defer row.Close()

	for row.Next() {
		var id, title, description, location, organizerName, organizerEmail, startTime, endTime string

		err := row.Scan(
			&id,
			&title,
			&description,
			&location,
			&organizerName,
			&organizerEmail,
			&startTime,
			&endTime,
		)
		if err != nil {
			return []models.Event{}, err
		}

		events = append(
			events,
			models.Event{
				Id:             id,
				Title:          title,
				Description:    description,
				Location:       location,
				OrganizerName:  organizerName,
				OrganizerEmail: organizerEmail,
				StartTime:      startTime,
				EndTime:        endTime,
			})
	}

	return events, nil
}

func (r *PostgreRepository) CreateEvent(
	ctx context.Context,
	event models.Event,
) (models.Event, error) {
	query := `insert into events(title, description, location, organizerName, organizerEmail, startTime, endTime) values($1, $2, $3, $4, $5, $6, $7) RETURNING *;`

	var id, title, description, location, organizerName, organizerEmail, startTime, endTime string
	err := r.db.QueryRowContext(
		ctx,
		query,
		event.Title,
		event.Description,
		event.Location,
		event.OrganizerName,
		event.OrganizerEmail,
		event.StartTime,
		event.EndTime,
	).Scan(
		&id,
		&title,
		&description,
		&location,
		&organizerName,
		&organizerEmail,
		&startTime,
		&endTime,
	)

	eventCreated := models.Event{
		Id:             id,
		Title:          title,
		Description:    description,
		Location:       location,
		OrganizerName:  organizerName,
		OrganizerEmail: organizerEmail,
		StartTime:      startTime,
		EndTime:        endTime,
	}

	if err != nil {
		return models.Event{}, err
	}

	return eventCreated, nil
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
	query := `update events set title=$1, description=$2, location=$3, organizerName=$4, organizerEmail=$5, startTime=$6, endTime=$7  where id=$8;`

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
		event.Id,
	)
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
