package main

import (
    "fmt"
    "log"
    "github.com/streadway/amqp"
    "product-management-system/async/image_processing/processor"
)

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %s", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %s", err)
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "image_processing_queue", // Queue name
        true,                     // durable
        false,                    // delete when unused
        false,                    // exclusive
        false,                    // no-wait
        nil,                      // arguments
    )
    if err != nil {
        log.Fatalf("Failed to declare a queue: %s", err)
    }

    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %s", err)
    }

    // Start consuming messages
    for msg := range msgs {
        imageURL := string(msg.Body)
        fmt.Printf("Received image URL: %s\n", imageURL)

        // Process the image (e.g., download, resize, etc.)
        processor.ProcessImage(imageURL)
    }
}
