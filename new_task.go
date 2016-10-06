package main

import (
	"log"
	"os"

	"github.com/olibob/amqptest/utilities"

	"github.com/streadway/amqp"
)

func main() {
	// connect to the RabbitMQ server
	conn, err := amqp.Dial("amqp://bob:bob*@192.168.60.31:5672/dev-vhost")
	utilities.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// create a channel
	ch, err := conn.Channel()
	utilities.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"taskQueue", // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utilities.FailOnError(err, "Failed to declare a queue")

	body := utilities.BodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	utilities.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
