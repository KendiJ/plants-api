package config

import (
    "os"
    "fmt"
)

type Config struct {
    DBUser     string
    DBPassword string
    DBName     string
    Port       string
}

func LoadConfig() (*Config, error) {
    // Simple validation example
    if os.Getenv("DB_PASSWORD") == "" {
        return nil, fmt.Errorf("DB_PASSWORD environment variable is required")
    }

    return &Config{
        DBUser:     getEnv("DB_USER", "root"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBName:     getEnv("DB_NAME", "homeplants"),
        Port:       getEnv("PORT", "8080"),
    }, nil
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}