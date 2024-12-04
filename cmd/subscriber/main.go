package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/internal/app/controller"
	"food-delivery-app-notification-service/internal/app/repository"
	"food-delivery-app-notification-service/internal/app/service"
	"food-delivery-app-notification-service/pkg/model"
	"food-delivery-app-notification-service/pkg/rabbitmq"
	"food-delivery-app-notification-service/server"
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
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	// Panic Recover Functionality
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered in subscriber: ", r)
		}
	}()

	dbConnection, _ := server.RDBMS(ctx)
	subscriberRepo := repository.NewRepoInit(dbConnection)
	rbtmq := rabbitmq.NewRabbitMQ(&config.Environment.MQ.ConString)
	subscriberService := service.NewSubscriberService(subscriberRepo)
	subscriberController := controller.NewSubscriberController(subscriberService)

	conn := rbtmq.OpenConnection()
	defer conn.Close()

	ch := rbtmq.CreateChannel(conn)
	defer ch.Close()
	if !conn.IsClosed() {
		var payload *model.MQPayload
		que := rbtmq.DeclareQueue(ch, &publisherQueName, false, false, false, false, nil)
		messages := rbtmq.ConsumeContent(ch, que)

		go func() {
			for d := range messages {
				err := json.Unmarshal(d.Body, &payload)
				if err != nil {
					log.Panic("Error while unmarshal message queue payload")
				}
				actualNotificationMessage := subscriberController.PrepairNotification(ctx, payload)
				subscriberController.SaveNotification(ctx, payload, actualNotificationMessage)
				subscriberController.SendNotification(ctx, actualNotificationMessage, payload.NotificationType)
			}
		}()

		log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
		<-forever
	} else {
		log.Panic("Message Queue is not alive")
	}
}
