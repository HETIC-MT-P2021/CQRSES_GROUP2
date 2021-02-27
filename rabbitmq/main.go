// RabbitMq Consumer
package main

import (
	"rabbitmq/consumer"
)

// main launch app consumer
func main() {
	consumer.ReceiveMessage()
}
