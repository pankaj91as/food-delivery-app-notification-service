package config

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

var Environment loader

type loader struct {
	APP application
	MQ  messageQue
}

type application struct {
	Host *string `env:"APP_HOST, default=0.0.0.0"`
	Port *string `env:"APP_PORT, default=8000"`
}

type messageQue struct {
	MQUsername *string `env:"MQ_USERNAME, default=admin"`
	MQPassword *string `env:"MQ_PASSWORD, default=admin"`
	MQHost     *string `env:"MQ_HOST, default=localhost"`
	MQPort     *string `env:"MQ_PORT, default=5672"`
	ConString  string
}

func init() {
	_, err := env.UnmarshalFromEnviron(&Environment)
	if err != nil {
		log.Fatal(err)
	}

	// Set RabbitMQ message string
	Environment.MQ.ConString = fmt.Sprintf("amqp://%s:%s@%s:%s/", *Environment.MQ.MQUsername, *Environment.MQ.MQPassword, *Environment.MQ.MQHost, *Environment.MQ.MQPort)
}
