package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
)

func Connect() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(config.RabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("Failed to open channel: %v", err)
	}
	return conn, ch
}
