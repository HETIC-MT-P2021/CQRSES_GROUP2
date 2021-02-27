package producer

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	numberOftries = 10
	waitTimes 		= 5
)

// SendMessage send messages to RabbitMq
func SendMessage() {
	// Connect to RabbitMq Instance
	var conn *amqp.Connection
	var err error
	for index := 0; index < numberOftries; index++ {
		conn, err = amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			time.Sleep(waitTimes * time.Second)
		} else {
			break
		}
	}

	// Open channel to RabbitMQ instance
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// Create RabbitMq queue
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)
	if err != nil {
		fmt.Println(err)
	}

	// Publish message in RabbitMq queue
	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
  fmt.Println("Successfully Published Message to Queue")
}
