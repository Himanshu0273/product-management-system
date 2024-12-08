package caching

import (
    "fmt"
    "log"
    "product-management-system/config"
    "github.com/go-redis/redis/v8"
    "context"
)

// RedisClient is the Redis client instance
var RedisClient *redis.Client

// InitRedis initializes the Redis connection using config values
func InitRedis(cfg *config.Config) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
    })

    // Test the Redis connection
    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }

    log.Println("Redis connection established.")
}
