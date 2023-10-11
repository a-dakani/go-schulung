package http

type Config struct {
	Host string `envconfig:"HOST" required:"true"`
	Port string `envconfig:"PORT" required:"true"`
}
