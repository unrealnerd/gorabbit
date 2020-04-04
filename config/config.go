package config

type Config struct {
	RabbitMQ struct {
		Host     string `yaml:"hosts"`
		Port     int    `yaml:"port"`
		UserName string    `yaml:"username"`
		Password string    `yaml:"password"`
		VHost string    `yaml:"vhost"`
	} `yaml:rabbitmq`
}
