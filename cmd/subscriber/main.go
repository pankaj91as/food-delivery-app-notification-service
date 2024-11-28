package main

import (
	"context"
	"flag"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/pkg/rabbitmq"
	"log"
	"time"
)

func main() {
	var publisherQueName string
	flag.StringVar(&publisherQueName, "queue-name", "priority", "Queue name should be pass to reterive messages eg: priority, pramotional")
	flag.Parse()

	// Define Required Variables
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

	ch := rbtmq.CreateChannel(conn)
	defer ch.Close()
	if !conn.IsClosed() {
		que := rbtmq.DeclareQueue(ch, &publisherQueName, false, false, false, false, nil)

		messages := rbtmq.ConsumeContent(ch, que)

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
