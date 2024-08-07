package events

import (
	"context"
	"events-manager/domain/events/models"
	configs "events-manager/infrastructure/configs/postgres"
	"events-manager/infrastructure/events/errors"
	"events-manager/pkgs/logger"
	"fmt"

	"golang.org/x/exp/slices"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreEventsRepository struct {
	db *gorm.DB
}

func NewPostgreEventsRepository(l logger.Logger, settings configs.PostgreSettigs) *PostgreEventsRepository {
	db, err := gorm.Open(postgres.Open(createConnToString(settings)), &gorm.Config{})
	if err != nil {
		l.Errorf("Error connecting [PostgreEventsRepository] to the DB: %s\n", err.Error())
	}

	db.AutoMigrate(&event{})

	return &PostgreEventsRepository{
		db,
	}
}

func (r *PostgreEventsRepository) GetEventById(
	ctx context.Context,
	id string,
) (models.Event, error) {
	var event event
	r.db.WithContext(ctx).Where("id = ?", id).First(&event)

	return *mapPostgresEventToEvent(event), nil
}

// This function is used internally to get the event by id
// without mapping it to the domain model
func (r *PostgreEventsRepository) getEventByIdInternal(
	ctx context.Context,
	id string,
) (event, error) {
	var event event
	r.db.WithContext(ctx).Where("id = ?", id).First(&event)

	return event, nil
}

func (r *PostgreEventsRepository) GetAllEvents(
	ctx context.Context,
) ([]models.Event, error) {
	var events []event
	r.db.WithContext(ctx).Find(&events)

	mappedEvents := make([]models.Event, len(events))
	for i, event := range events {
		mappedEvents[i] = *mapPostgresEventToEvent(event)
	}

	return mappedEvents, nil
}

func (r *PostgreEventsRepository) CreateEvent(
	ctx context.Context,
	event models.Event,
) (models.Event, error) {
	postgresEvent := mapEventToPostgresEvent(event)
	r.db.WithContext(ctx).Create(postgresEvent)

	event.Id = fmt.Sprint(postgresEvent.ID)

	return event, nil
}

func (r *PostgreEventsRepository) AddAttendeToEventById(
	ctx context.Context,
	id string,
	attendeeEmail string,
) ([]string, error) {
	event, err := r.getEventByIdInternal(ctx, id)
	if err != nil {
		return []string{}, err
	}
	if slices.Contains(event.Attendees, attendeeEmail) {
		return event.Attendees, errors.NewDuplicateAttendeeError()
	}

	event.Attendees = append(event.Attendees, attendeeEmail)
	r.db.WithContext(ctx).Model(&event).Update(
		"attendees",
		gorm.Expr("array_append(attendees, ?)", attendeeEmail),
	)

	return event.Attendees, nil
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
	postgresEvent := mapEventToPostgresEvent(event)
	r.db.WithContext(ctx).Save(postgresEvent)

	return event, nil
}

func (r *PostgreEventsRepository) DeleteEventById(
	ctx context.Context,
	id string,
) (models.Event, error) {
	r.db.WithContext(ctx).Delete(&event{}, id)
	return models.Event{}, nil
}

// Take our connection struct and convert to a string for our db connection info
func createConnToString(info configs.PostgreSettigs) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.DBName)
}
