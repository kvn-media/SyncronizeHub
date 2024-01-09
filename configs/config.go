// config/config.go

package config

import (
	"fmt"
	"os"
)

// Config holds the application configuration
type Config struct {
   DBUsername     string
   DBPassword     string
   DBName         string
   DBSSLMode      string
   // Add other configuration parameters as needed
}

// NewConfig creates a new Config instance with values from environment variables
func NewConfig() *Config {
	return &Config{
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE"),
		// Initialize other configuration parameters
	}
}

// LoadConfigFromEnv loads configuration from environment variables
func LoadConfigFromEnv(cfg *Config) {
	cfg.DBUsername = os.Getenv("DB_USERNAME")
	cfg.DBPassword = os.Getenv("DB_PASSWORD")
	cfg.DBName = os.Getenv("DB_NAME")
	cfg.DBSSLMode = os.Getenv("DB_SSL_MODE")
	// Load other configuration parameters as needed
}

// Validate checks if the configuration is valid
func (cfg *Config) Validate() error {
	if cfg.DBUsername == "" || cfg.DBPassword == "" || cfg.DBName == "" || cfg.DBSSLMode == "" {
		return fmt.Errorf("Incomplete configuration. Please provide all required parameters.")
	}
	// Validate other configuration parameters if needed
	return nil
}
