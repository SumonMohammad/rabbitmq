package producer

import (
	"log"
	"rabbitmq-app/internal/rabbitmq"
)

func Send(ch *rabbitmq.ChannelWrapper, body string) {
	// we will call internal publisher directly (ch is *amqp.Channel)
	err := rabbitmq.PublishImmediate(ch.Channel, body)
	if err != nil {
		log.Printf("Failed to send immediate message: %v", err)
		return
	}
	log.Printf("Sent immediate: %s", body)
}
