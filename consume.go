package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	// connect to RabbitMQ
	conn, err := amqp.Dial("amqp://bob:bob*@192.168.60.31:5672/dev-vhost")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channle")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // Name
		false,   // durable
		false,   // delete when used
		false,   // exclusice
		false,   // no-wait
		nil,     // arguments
	)

	failOnError(err, "Failed to declare queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)

	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
		}
	}()

	log.Printf(" [x] Waiting for messages. To exit press CTRL-C")
	<-forever
}
