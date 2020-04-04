package helper

import(
	"log"
	"fmt"
	"net/url"
	"strconv"

	"github.com/spf13/viper"

	c "gorabbit/config"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func GetRabbitMQURL() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var config c.Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Host is\t", config.RabbitMQ.Host)
	fmt.Println("UserName is\t", url.PathEscape(config.RabbitMQ.UserName))

	rabbitMQURL := "amqp://" +
		url.PathEscape(config.RabbitMQ.UserName) +
		":" +
		url.PathEscape(config.RabbitMQ.Password) +
		"@" +
		config.RabbitMQ.Host +
		":" +
		strconv.Itoa(config.RabbitMQ.Port) +
		"/" +
		config.RabbitMQ.VHost

	return rabbitMQURL
}