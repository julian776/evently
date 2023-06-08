package reminders

import (
	"context"
	"fmt"
	"notifier/domain/reminders/models"
	"notifier/pkgs/logger"
	"time"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RemindersMongoRepository struct {
	logger       logger.Logger
	mongoSettigs MongoSettigs
	db           *mongo.Database
	coll         *mongo.Collection
}

func NewRemindersMongoRepository(l logger.Logger, settings MongoSettigs) *RemindersMongoRepository {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(settings.Url))
	if err != nil {
		l.Errorf("can not connect to mongo %s", err.Error())
	}
	db := client.Database(settings.DBName)
	coll := db.Collection(settings.CollName)

	return &RemindersMongoRepository{
		l,
		settings,
		db,
		coll,
	}
}

func (r *RemindersMongoRepository) CreateReminder(
	ctx context.Context,
	reminder models.Reminder,
) (models.Reminder, error) {
	var reminderCreated models.Reminder

	res, err := r.coll.InsertOne(
		ctx,
		reminder,
	)
	if err != nil {
		return reminderCreated, fmt.Errorf("can not create reminder %s", err.Error())
	}
	mapstructure.Decode(res, &reminderCreated)
	return reminderCreated, nil
}

func (r *RemindersMongoRepository) GetAllTodayReminders(ctx context.Context) ([]models.Reminder, error) {
	return []models.Reminder{}, nil
}
func (r *RemindersMongoRepository) AddNewEmailToReminderWithEventId(ctx context.Context, eventId string, email string) error {
	return nil
}

func (r *RemindersMongoRepository) DeletePastReminders(ctx context.Context) error {
	return nil
}
