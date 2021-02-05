package consumer

import (
	"fmt"

	"github.com/streadway/amqp"
)

// ReceiveMessage receive event from rabbitmq
func ReceiveMessage() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	// Open channel to RabbitMQ instance
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer channel.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := channel.Consume("TestQueue", "", true, false, false, false, nil)

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
}
