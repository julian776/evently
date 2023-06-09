package emails

type Settings struct {
	Email    string `envconfig:"EMAIL" default:""`
	Password string `envconfig:"EMAIL_PASSWORD" default:""`
}
