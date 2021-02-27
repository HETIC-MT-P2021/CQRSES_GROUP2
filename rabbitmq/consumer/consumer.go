package consumer

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

const (
	numberOftries = 10
	waitTimes  		= 5
)

// ReceiveMessage receive messages from RabbitMq
func ReceiveMessage() {
	// Connect to RabbitMq instance
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

	if err != nil {
		fmt.Println(err)
	}

	// Consume RabbitMq queue
	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// Get RabbitMq messages
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()
}
