package handler

import (
	"context"
	"fmt"
	"food-delivery-app-notification-service/internal/app/config"
	"food-delivery-app-notification-service/pkg/model"
	"food-delivery-app-notification-service/pkg/rabbitmq"
	"log"
	"time"
)

func Publish(publisherQueName string, messageBody *model.MQPayload) {
	// Define Context with timeout 5 Second
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
	if !conn.IsClosed() {
		ch := rbtmq.CreateChannel(conn)
		defer ch.Close()

		que := rbtmq.DeclareQueue(ch, &messageBody.QueueName, false, false, false, false, nil)

		err := rbtmq.PublishContent(ch, que, ctx, "", false, false, messageBody.Message)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		log.Panic("Message Queue is not alive")
	}

	log.Printf(" [%s] Sent %s in %s\n", messageBody.NotificationType, messageBody.Message, messageBody.QueueName)
}
