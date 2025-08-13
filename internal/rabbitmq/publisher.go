package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
)

func PublishImmediate(ch *amqp.Channel, body string) error {
	return ch.PublishWithContext(context.Background(),
		config.ExchangeMain,
		config.RoutingKey,
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func PublishWithDelay(ch *amqp.Channel, body string, delayMs int) error {
	headers := amqp.Table{"x-delay": int32(delayMs)}
	return ch.PublishWithContext(context.Background(),
		config.ExchangeDelayed,
		config.RoutingKey, // delayed exchange will re-publish to bound queues with same routing key
		false, false,
		amqp.Publishing{
			Headers:     headers,
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func PublishToDead(ch *amqp.Channel, body string) error {
	return ch.PublishWithContext(context.Background(),
		config.ExchangeDead,
		"dead",
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
}

func mustPublish(publishErr error) {
	if publishErr != nil {
		log.Fatalf("publish failed: %v", publishErr)
	}
}
