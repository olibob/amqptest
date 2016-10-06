package main

import (
	"bytes"
	"log"
	"time"

	"github.com/olibob/amqptest/utilities"
	"github.com/streadway/amqp"
)

func main() {
	// connect to RabbitMQ
	conn, err := amqp.Dial("amqp://bob:bob*@192.168.60.31:5672/dev-vhost")
	utilities.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// open channel
	ch, err := conn.Channel()
	utilities.FailOnError(err, "Failed to open channle")
	defer ch.Close()
	// idempotent
	q, err := ch.QueueDeclare(
		"taskQueue", // Name
		true,        // durable
		false,       // delete when used
		false,       // exclusice
		false,       // no-wait
		nil,         // arguments
	)

	utilities.FailOnError(err, "Failed to declare queue")

	err = ch.Qos(
		2,     // prefetch Count
		0,     // prefetch size
		false, // global
	)

	utilities.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)

	utilities.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf(" [x] Waiting for messages. To exit press CTRL-C")
	<-forever
}
