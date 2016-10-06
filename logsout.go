package main

import (
	"log"

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
		"",    // Name
		false, // durable
		false, // delete when used
		true,  // exclusice
		false, // no-wait
		nil,   // arguments
	)

	utilities.FailOnError(err, "Failed to declare queue")

	err = ch.QueueBind(
		q.Name, //name
		"",     //key
		"logs", //exchange
		false,  //noWait
		nil,    //args
	)

	utilities.FailOnError(err, "Failed to create binding")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	utilities.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Recieved message: %s", d.Body)
		}
	}()

	log.Printf(" [x] Waiting for messages. To exit press CTRL-C")
	<-forever
}
