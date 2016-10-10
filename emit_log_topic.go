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

	err = ch.ExchangeDeclare(
		"logsTopic", //name
		"topic",     //kind
		true,        //durable
		false,       //autoDelete
		false,       //internal
		false,       //noWait
		nil,         //args
	)
	utilities.FailOnError(err, "Failed to declare an exchange")

	body := utilities.BodyFrom2(os.Args)
	err = ch.Publish(
		"logsTopic",                     // exchange
		utilities.SeverityFrom(os.Args), // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})

	utilities.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}
