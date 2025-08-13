package consumer

import (
	"log"
	//"github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
	"rabbitmq-app/internal/rabbitmq"
)

func Start(ch *rabbitmq.ChannelWrapper) {
	err := ch.Channel.Qos(1, 0, false) // 1 message at a time
	if err != nil {
		log.Fatalf("Failed to set QoS: %v", err)
	}

	msgs, err := ch.Channel.Consume(
		config.MainQueue,
		"",
		false,
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
		HandleDelivery(ch.Channel, d)
	}
}
