package users

type UsersSettings struct {
	Queue string `envconfig:"USERS_QUEUE" default:"users"`
}
