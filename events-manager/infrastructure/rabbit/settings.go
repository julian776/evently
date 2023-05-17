package rabbit

type Settings struct {
	Url   string `envconfig:"RABBIT_URL" required:"true"`
	Queue string `envconfig:"RABBIT_QUEUE" default:"events"`
}
