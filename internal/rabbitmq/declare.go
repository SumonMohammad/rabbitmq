package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"rabbitmq-app/config"
)

func DeclareAll(ch *amqp.Channel) {
	// 1) main direct exchange
	err := ch.ExchangeDeclare(
		config.ExchangeMain,
		"direct",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("ExchangeDeclare main failed: %v", err)
	}

	// 2) dead exchange
	err = ch.ExchangeDeclare(
		config.ExchangeDead,
		"direct",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("ExchangeDeclare dead failed: %v", err)
	}

	// 3) delayed exchange (plugin)
	args := amqp.Table{
		"x-delayed-type": "direct",
	}
	err = ch.ExchangeDeclare(
		config.ExchangeDelayed,
		"x-delayed-message",
		true, false, false, false, args,
	)
	if err != nil {
		log.Fatalf("ExchangeDeclare delayed failed: %v", err)
	}

	// 4) main queue
	_, err = ch.QueueDeclare(
		config.MainQueue,
		true,  false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("QueueDeclare main failed: %v", err)
	}

	// 5) dead queue
	_, err = ch.QueueDeclare(
		config.DeadQueue,
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("QueueDeclare dead failed: %v", err)
	}

	// 6) Bindings
	if err := ch.QueueBind(config.MainQueue, config.RoutingKey, config.ExchangeMain, false, nil); err != nil {
		log.Fatalf("QueueBind main failed: %v", err)
	}
	if err := ch.QueueBind(config.DeadQueue, "dead", config.ExchangeDead, false, nil); err != nil {
		log.Fatalf("QueueBind dead failed: %v", err)
	}

	log.Println("Declared exchanges & queues (main, delayed, dead) successfully")
}
