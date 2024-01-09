package utils

import (
	"github.com/google/uuid"
)

// GenerateID generates a new UUID as a string.
func GenerateID() string {
	return uuid.New().String()
}
