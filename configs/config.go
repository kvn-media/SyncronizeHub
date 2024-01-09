// config/config.go

package config

import "os"

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
      DBUsername:     os.Getenv("DB_USERNAME"),
      DBPassword:     os.Getenv("DB_PASSWORD"),
      DBName:         os.Getenv("DB_NAME"),
      DBSSLMode:      os.Getenv("DB_SSL_MODE"),
      // Initialize other configuration parameters
   }
}
