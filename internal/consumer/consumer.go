package consumer

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
)

func Start(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		config.MainQueue,
		"",
		false, // manual ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}

	log.Println("Consumer started, waiting for messages...")
	for d := range msgs {
		HandleDelivery(ch, d)
	}
}
