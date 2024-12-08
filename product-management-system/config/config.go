package config

import (
    "log"
    "os"

    "github.com/spf13/viper"
)

// Config holds the configuration for database, redis, rabbitmq, and log settings
type Config struct {
    Database struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        User     string `mapstructure:"user"`
        Password string `mapstructure:"password"`
        DBName   string `mapstructure:"dbname"`
        SSLMode  string `mapstructure:"sslmode"`
    } `mapstructure:"database"`

    Redis struct {
        Host string `mapstructure:"host"`
        Port int    `mapstructure:"port"`
    } `mapstructure:"redis"`

    RabbitMQ struct {
        Host     string `mapstructure:"host"`
        Port     int    `mapstructure:"port"`
        Username string `mapstructure:"username"`
        Password string `mapstructure:"password"`
    } `mapstructure:"rabbitmq"`

    Log struct {
        Level string `mapstructure:"level"`
    } `mapstructure:"log"`
}

// LoadConfig loads the configuration from the config.yaml file
func LoadConfig() (*Config, error) {
    viper.SetConfigName("config")     // Name of the config file (without extension)
    viper.SetConfigType("yaml")       // Set the type to YAML
    viper.AddConfigPath("./config/")  // Path to look for the config file

    // Read the config file
    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    // Unmarshal the config into the Config struct
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}
