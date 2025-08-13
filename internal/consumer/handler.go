package consumer

import (
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
	"rabbitmq-app/internal/rabbitmq"
)

// getRetryFromHeaders reads x-retry header if present
func getRetryCount(headers amqp.Table) int {
	if headers == nil {
		return 0
	}
	if v, ok := headers["x-retry"]; ok {
		switch t := v.(type) {
		case int:
			return t
		case int32:
			return int(t)
		case int64:
			return int(t)
		case string:
			i, _ := strconv.Atoi(t)
			return i
		}
	}
	return 0
}

// HandleDelivery processes a message; on failure it schedules retry or dead
func HandleDelivery(ch *amqp.Channel, d amqp.Delivery) {
	body := string(d.Body)
	log.Printf("Processing message: %s", body)

	// simplistic simulated failure: if body == "fail" treat as failure
	if body == "fail" {
		// read retry count
		attempt := getRetryCount(d.Headers)
		if attempt >= config.MaxRetries {
			log.Printf("Max retries reached (%d). Sending to dead queue.", attempt)
			if err := rabbitmq.PublishToDead(ch, body); err != nil {
				log.Printf("Failed to publish to dead: %v", err)
			}
			d.Ack(false)
			return
		}

		// increment and publish to delayed exchange (e.g., 5s * (attempt+1))
		newAttempt := attempt + 1
		delayMs := 5000 * newAttempt // exponential-ish backoff
		headers := amqp.Table{"x-retry": newAttempt, "x-delay": int32(delayMs)}

		err := ch.Publish(
			config.ExchangeDelayed,
			config.RoutingKey,
			false, false,
			amqp.Publishing{
				Headers:     headers,
				ContentType: "text/plain",
				Body:        d.Body,
			},
		)
		if err != nil {
			log.Printf("Failed to publish to delayed exchange: %v", err)
			// if publish to delayed fails, publish to dead to avoid infinite loops
			_ = rabbitmq.PublishToDead(ch, body)
			d.Ack(false)
			return
		}
		log.Printf("Message requeued to delayed exchange with delay %d ms (attempt %d).", delayMs, newAttempt)
		d.Ack(false)
		return
	}

	// success path
	log.Printf("Processed successfully: %s", body)
	d.Ack(false)
}
