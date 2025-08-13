package main

import (
	"log"
	"os"
	"rabbitmq-app/internal/rabbitmq"
	//"rabbitmq-app/config"
)

func main() {
	// override RabbitURL for docker-compose: if you want docker internal host use amqp://guest:guest@rabbitmq:5672/
	// config.RabbitURL = "amqp://guest:guest@rabbitmq:5672/"

	conn, ch := rabbitmq.Connect()
	defer conn.Close()
	defer ch.Close()

	// ensure exchanges & queues declared
	rabbitmq.DeclareAll(ch)

	var message string
	if len(os.Args) > 1 {
		message = os.Args[1]
	} else {
		message = "hello"
	}

	err := rabbitmq.PublishImmediate(ch, message)
	if err != nil {
		log.Fatalf("PublishImmediate failed: %v", err)
	}
	log.Println("Published immediate message:", message)
}
