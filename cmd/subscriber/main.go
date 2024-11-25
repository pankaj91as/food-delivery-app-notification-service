package main

import (
	"context"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/pkg/rabbitmq"
	"log"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// Define Required Variables
	publisherQueName := "hello"
	var forever chan struct{}

	// Define Context with timeout 5 Second
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Panic Recover Functionality
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered in publisher: ", r)
		}
	}()

	rbtmq := rabbitmq.NewRabbitMQ(&config.Environment.MQ.ConString)

	conn := rbtmq.OpenConnection()
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()
	if !conn.IsClosed() {
		que := rbtmq.DeclareQueue(ch, &publisherQueName, false, false, false, false, nil)

		messages := rbtmq.ConsumeContent(ch, que)
		if err != nil {
			fmt.Println(err)
		}

		go func() {
			for d := range messages {
				log.Printf("Received a message: %s", d.Body)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	} else {
		log.Panic("Message Queue is not alive")
	}
}
