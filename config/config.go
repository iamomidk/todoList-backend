// config/config.go

package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration.
type Config struct {
	DatabaseURL string
}

// NewConfig creates a new instance of Config by reading from environment variables.
func NewConfig() (*Config, error) {
	dbURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	return &Config{
		DatabaseURL: dbURL,
	}, nil
}
