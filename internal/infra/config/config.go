package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config represents application-wide configuration
type Config struct {
	DatabaseURL string `json:"database_url"`
	ServerPort  string `json:"server_port"`
	SecretKey   string `json:"secret_key"`
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	// Unmarshal configuration from environment variables
	SERVER_PORT := os.Getenv("SERVER_PORT")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	DATABSE_URL := os.Getenv("DATABASE_URL")

	config := &Config{
		DatabaseURL: DATABSE_URL,
		ServerPort:  SERVER_PORT,
		SecretKey:   SECRET_KEY,
	}

	// Validate required configuration values
	required := []string{"DATABASE_URL", "SERVER_PORT", "SECRET_KEY"}
	for _, key := range required {
		if value := os.Getenv(key); value == "" {
			return nil, fmt.Errorf("%s environment variable is required", key)
		}
	}

	return config, nil
}
