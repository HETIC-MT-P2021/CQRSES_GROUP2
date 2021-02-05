// RabbitMQ
package main

import (
	"rabbitmq/consumer"
	"rabbitmq/producer"
)

// main launch RabbitMQ message broker
func main() {
	producer.SendMessage()
	consumer.ReceiveMessage()
}
