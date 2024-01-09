// models/user.go

package models

import (
	"time"
)

// UserModel represents a user in the system.
type UserModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUserModel creates a new instance of UserModel with the given parameters.
func NewUserModel(name, email, password string) *UserModel {
	return &UserModel{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Validate method for UserModel to perform validation checks.
func (u *UserModel) Validate() error {
	// Perform validation checks, e.g., check if email is valid, password meets requirements, etc.
	// Return an error if validation fails.
	return nil
}
