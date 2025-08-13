package consumer

import (
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

func getRetryCount(headers amqp.Table) int {
	if headers == nil {
		return 0
	}
	if v, ok := headers["x-retry"]; ok {
		switch t := v.(type) {
		case int:
			return t
		case int8:
			return int(t)
		case int16:
			return int(t)
		case int32:
			return int(t)
		case int64:
			return int(t)
		case uint8:
			return int(t)
		case uint16:
			return int(t)
		case uint32:
			return int(t)
		case uint64:
			return int(t)
		case float32:
			return int(t)
		case float64:
			return int(t)
		case string:
			i, err := strconv.Atoi(t)
			if err != nil {
				return 0
			}
			return i
		default:
			log.Printf("Unknown type for x-retry: %v", t)
			return 0
		}
	}
	return 0
}

func HandleDelivery(ch *amqp.Channel, d amqp.Delivery) {
	retryCount := getRetryCount(d.Headers)
	if retryCount >= 3 { // Max retries
		log.Printf("Max retries reached (%d). Sending to dead queue.", retryCount)
		ch.Publish(
			"dead_exchange", // dead letter exchange
			"",              // routing key
			false,           // mandatory
			false,           // immediate
			amqp.Publishing{
				ContentType: d.ContentType,
				Body:        d.Body,
			},
		)
		d.Ack(false) // Acknowledge and remove from queue
		return
	}

	// Simulate failure
	log.Printf("Processing message: %s (attempt: %d)", string(d.Body), retryCount)
	newHeaders := make(amqp.Table)
	if retryCount > 0 {
		newHeaders["x-retry"] = retryCount + 1
	} else {
		newHeaders["x-retry"] = 1
	}
	newHeaders["x-delay"] = 5000 * (retryCount + 1) // Increase delay per attempt (5s, 10s, 15s)

	ch.Publish(
		"delayed_exchange", // delayed exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: d.ContentType,
			Body:        d.Body,
			Headers:     newHeaders,
		},
	)
	d.Ack(false) // Acknowledge after requeuing
}
