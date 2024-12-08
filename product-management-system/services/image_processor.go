package services

import (
    "log"
    "github.com/streadway/amqp"
)

func ConsumeQueue() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
    }
    defer ch.Close()

    msgs, err := ch.Consume("image_queue", "", true, false, false, false, nil)
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    for msg := range msgs {
        log.Printf("Processing image: %s", msg.Body)
    }
}
