package rabbitmq

import (
	"context"
	"log"

	"rabbitmq-app/config"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishImmediate(ch *amqp.Channel, body string) error {
	headers := amqp.Table{"x-retry": 0}
	return ch.PublishWithContext(context.Background(),
		config.ExchangeMain,
		config.RoutingKey,
		false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
			Headers:     headers,
		},
	)
}

func PublishWithDelay(ch *amqp.Channel, body string, delayMs int) error {
	headers := amqp.Table{"x-delay": int32(delayMs)}
	return ch.PublishWithContext(context.Background(),
		config.ExchangeDelayed,
		config.RoutingKey,
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
