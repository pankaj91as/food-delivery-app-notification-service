package rabbitmq

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQ interface {
	OpenConnection() *amqp.Connection
	CreateChannel(conn *amqp.Connection) *amqp.Channel
	DeclareQueue(ch *amqp.Channel, QueueName *string, Durable, DeleteUnUsed, Exclusive, NoWait bool, arg amqp.Table) amqp.Queue
	PublishContent(ch *amqp.Channel, queue amqp.Queue, ctx context.Context, exchange string, mandatory, immediate bool, body string) error
	ConsumeContent(ch *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery
}
type RabbitMQ struct {
	endpoint *string
}

func NewRabbitMQ(conString *string) IRabbitMQ {
	return &RabbitMQ{
		endpoint: conString,
	}
}

func (rmq *RabbitMQ) OpenConnection() *amqp.Connection {
	fmt.Println("RabbitMQ Listening ON:", *rmq.endpoint)

	// Try to open RabbitMQ connection
	conn, err := amqp.Dial(*rmq.endpoint)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	// send connection back
	return conn
}

func (rmq *RabbitMQ) CreateChannel(conn *amqp.Connection) *amqp.Channel {
	fmt.Println("RabbitMQ Connection State:", !conn.IsClosed())

	ch, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a channel", err)
	}

	// send channel back
	return ch
}

func (rmq *RabbitMQ) DeclareQueue(ch *amqp.Channel, QueueName *string, Durable, DeleteUnUsed, Exclusive, NoWait bool, arg amqp.Table) amqp.Queue {
	fmt.Println("RabbitMQ Channel State:", !ch.IsClosed())

	q, err := ch.QueueDeclare(
		*QueueName,   // name
		Durable,      // durable
		DeleteUnUsed, // delete when unused
		Exclusive,    // exclusive
		NoWait,       // no-wait
		arg,          // arguments
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to declare a queue", err)
	}

	// send queue back
	return q
}

func (rmq *RabbitMQ) PublishContent(ch *amqp.Channel, queue amqp.Queue, ctx context.Context, exchange string, mandatory, immediate bool, body string) error {
	fmt.Printf("Sending Message in queue: %s\n", queue.Name)

	err := ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		log.Panicf("%s: %s", "Failed to publish a message", err)
	}

	return err
}

func (rmq *RabbitMQ) ConsumeContent(ch *amqp.Channel, queue amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Panicf("%s: %s", "Failed to register a consumer", err)
	}

	return msgs
}
