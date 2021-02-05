package producer

import (
	"fmt"

	"github.com/streadway/amqp"
)

// SendMessage send event to rabbitmq
func SendMessage() {
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

	queue, err := channel.QueueDeclare("TestQueue", false, false, false, false, nil)
	fmt.Println(queue)

	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish message to queue!
	err = channel.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")
}
