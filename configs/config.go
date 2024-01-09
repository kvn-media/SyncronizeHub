// configs/config.go
package configs

import (
	"fmt"
	"os"
	"strconv"
)

// Config represents the application configuration.
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBDriver   string
	APIHost    string
	APIPort    int
	AppName    string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvAsInt("DB_PORT", 5432),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		APIHost:    getEnv("API_HOST", "localhost"),
		APIPort:    getEnvAsInt("API_PORT", 8080),
		AppName:    getEnv("APP_NAME", ""),
	}, nil
}

// getEnv retrieves the value of an environment variable with a default fallback.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// getEnvAsInt retrieves the value of an environment variable as an integer with a default fallback.
func getEnvAsInt(key string, fallback int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return fallback
}

// ValidateConfig validates the configuration and returns an error if any required field is missing.
func (c *Config) ValidateConfig() error {
	requiredFields := []string{"DBHost", "DBPort", "DBUser", "DBPassword", "DBName", "DBDriver", "APIHost", "APIPort", "AppName"}

	for _, field := range requiredFields {
		if value := getFieldValue(field, c); value == "" {
			return fmt.Errorf("missing required configuration field: %s", field)
		}
	}

	return nil
}

// getFieldValue retrieves the value of a field from the Config struct.
func getFieldValue(fieldName string, config *Config) string {
	switch fieldName {
	case "DBHost":
		return config.DBHost
	case "DBUser":
		return config.DBUser
	case "DBPassword":
		return config.DBPassword
	case "DBName":
		return config.DBName
	case "DBDriver":
		return config.DBDriver
	case "APIHost":
		return config.APIHost
	case "AppName":
		return config.AppName
	default:
		return ""
	}
}
