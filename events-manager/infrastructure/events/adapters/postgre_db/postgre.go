package postgredb

import (
	"context"
	"database/sql"
	"events-manager/domain/events/models"
	configs "events-manager/infrastructure/configs/postgres"
	"events-manager/pkgs/logger"
	"fmt"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type PostgreEventsRepository struct {
	db *sql.DB
}

func NewPostgreEventsRepository(l logger.Logger, settings configs.PostgreSettigs) *PostgreEventsRepository {
	fmt.Println(settings)
	db, err := sql.Open("postgres", createConnToString(settings))
	if err != nil {
		l.Errorf("Error connecting [PostgreEventsRepository] to the DB: %s\n", err.Error())
	}

	// check if we can ping our DB
	err = db.Ping()
	if err != nil {
		l.Errorf("Error [PostgreEventsRepository] could not ping database: %s\n", err.Error())
	}

	return &PostgreEventsRepository{
		db,
	}
}

func (r *PostgreEventsRepository) GetEventById(
	ctx context.Context,
	id string,
) (models.Event, error) {
	var event models.Event

	query := `select
id,
title,
description,
cost,
location,
attendees,
organizerName,
organizerEmail,
startTime,
endTime FROM events WHERE id=$1`

	row, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return event, err
	}

	defer row.Close()

	if row.Next() {
		var id, title, description, location, organizerName, organizerEmail, startTime, endTime string
		var attendees []string
		var cost float32
		err := row.Scan(
			&id,
			&title,
			&description,
			&cost,
			&location,
			pq.Array(&attendees),
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
			Cost:           cost,
			Location:       location,
			OrganizerName:  organizerName,
			OrganizerEmail: organizerEmail,
			StartTime:      startTime,
			EndTime:        endTime,
		}
	}

	return event, nil
}

func (r *PostgreEventsRepository) GetAllEvents(
	ctx context.Context,
) ([]models.Event, error) {
	var events []models.Event

	query := `select id, title, description, cost, location, attendees, organizerName, organizerEmail, startTime, endTime from events`
	row, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return []models.Event{}, err
	}

	defer row.Close()

	for row.Next() {
		var id, title, description, location, organizerName, organizerEmail, startTime, endTime string
		var attendees []string
		var cost float32
		err := row.Scan(
			&id,
			&title,
			&description,
			&cost,
			&location,
			pq.Array(&attendees),
			&organizerName,
			&organizerEmail,
			&startTime,
			&endTime,
		)
		if err != nil {
			return []models.Event{}, err
		}

		events = append(events, models.Event{
			Id:             id,
			Title:          title,
			Description:    description,
			Cost:           cost,
			Location:       location,
			Attendees:      attendees,
			OrganizerName:  organizerName,
			OrganizerEmail: organizerEmail,
			StartTime:      startTime,
			EndTime:        endTime,
		})
	}

	return events, nil
}

func (r *PostgreEventsRepository) CreateEvent(
	ctx context.Context,
	event models.Event,
) (models.Event, error) {
	query := `insert into events(
title,
description,
cost,
location,
attendees,
organizerName,
organizerEmail,
startTime,
endTime) values($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;`

	var id, title, description, location, organizerName, organizerEmail, startTime, endTime string
	var attendees []string
	var cost float32
	err := r.db.QueryRowContext(
		ctx,
		query,
		event.Title,
		event.Description,
		event.Cost,
		event.Location,
		pq.Array(event.Attendees),
		event.OrganizerName,
		event.OrganizerEmail,
		event.StartTime,
		event.EndTime,
	).Scan(
		&id,
		&title,
		&description,
		&cost,
		&location,
		pq.Array(&attendees),
		&organizerName,
		&organizerEmail,
		&startTime,
		&endTime,
	)

	eventCreated := models.Event{
		Id:             id,
		Title:          title,
		Description:    description,
		Cost:           cost,
		Location:       location,
		Attendees:      attendees,
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

func (r *PostgreEventsRepository) AddAttendeToEventById(
	ctx context.Context,
	id string,
	attendeeEmail string,
) ([]string, error) {
	query := `UPDATE events SET
attendees = array_append(attendees, $2)
WHERE id=$1 RETURNING attendees;`

	var attendees []string
	err := r.db.QueryRowContext(
		ctx,
		query,
		id,
		attendeeEmail,
	).Scan(
		pq.Array(&attendees),
	)

	if err != nil {
		return []string{}, err
	}

	return attendees, nil
}

func (r *PostgreEventsRepository) GetAllEventsByOrganizerEmail(
	ctx context.Context,
	id string,
) ([]models.Event, error) {
	return []models.Event{}, nil
}

func (r *PostgreEventsRepository) UpdateEvent(
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

func (r *PostgreEventsRepository) DeleteEventById(
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
func createConnToString(info configs.PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
