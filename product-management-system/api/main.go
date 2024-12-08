package main

import (
    "fmt"
    "log"
    "net/http"
    "product-management-system/api/router" // Correct import for router
    "product-management-system/config"    // Correct import for config
)

func main() {
    // Load the configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Example: Accessing the database configuration
    fmt.Println("Database Host:", cfg.Database.Host)
    fmt.Println("Database Port:", cfg.Database.Port)

    // Example: Accessing the Redis configuration
    fmt.Println("Redis Host:", cfg.Redis.Host)

    // Example: Accessing the RabbitMQ configuration
    fmt.Println("RabbitMQ Host:", cfg.RabbitMQ.Host)

    // Initialize the router and API routes
    r := router.SetupRouter()

    // Start the server on port 8080
    log.Println("Starting API server on port 8080...")
    err = http.ListenAndServe(":8080", r)
    if err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}
