package database

import (
    "fmt"
    "log"
    "product-management-system/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// InitDB initializes the database connection using config values
func InitDB(cfg *config.Config) {
    var err error

    // Use the database configuration from config.yaml
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
        cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.Port, cfg.Database.SSLMode)

    // Connect to the database using GORM
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    log.Println("Database connection established.")
}
