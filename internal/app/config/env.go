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
	Host *string `env:"APP_HOST"`
	Port *int    `env:"APP_PORT"`
}

type messageQue struct {
	MQUsername *string `env:"MQ_USERNAME"`
	MQPassword *string `env:"MQ_PASSWORD"`
	MQHost     *string `env:"MQ_HOST"`
	MQPort     *string `env:"MQ_PORT"`
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
