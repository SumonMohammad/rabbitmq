package main

import (
    "log"
    "os"
    "rabbitmq-app/internal/rabbitmq"
    "rabbitmq-app/internal/producer"
)

func main() {
    conn, ch := rabbitmq.Connect()
    defer conn.Close()
    defer ch.Channel.Close()

    rabbitmq.DeclareAll(ch.Channel)

    var message string
    if len(os.Args) > 1 {
        message = os.Args[1]
    } else {
        message = "hello"
    }

    producer.Send(ch, message)
    log.Println("Published immediate message:", message)
}