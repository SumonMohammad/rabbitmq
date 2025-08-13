package main

import (
	"log"
	"rabbitmq-app/internal/rabbitmq"
	"rabbitmq-app/internal/consumer"
)

func main() {
	conn, ch := rabbitmq.Connect()
	defer conn.Close()
	defer ch.Channel.Close()

	rabbitmq.DeclareAll(ch.Channel)

	log.Println("Starting consumer ...")
	consumer.Start(ch)

	// never exits
	select {}
}