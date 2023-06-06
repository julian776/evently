package reminders

type MongoSettigs struct {
	Url      string `envconfig:"MONGO_URL" required:"true"`
	DBName   string `envconfig:"MONGO_DB_NAME" required:"true"`
	CollName string `envconfig:"MONGO_COLL_NAME" required:"true"`
}
