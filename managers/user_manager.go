// managers/user_manager.go

package managers

import (
	"errors"

	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/repository"
	"golang.org/x/crypto/bcrypt"
)

// ErrInvalidCredentials is returned when login credentials are invalid
var ErrInvalidCredentials = errors.New("invalid credentials")

// UserManager handles user-related business logic
type UserManager struct {
	UserRepository repository.UserRepository
}

// NewUserManager creates a new instance of UserManager
func NewUserManager(userRepository repository.UserRepository) *UserManager {
	return &UserManager{
		UserRepository: userRepository,
	}
}

// RegisterUser handles the user registration process
func (um *UserManager) RegisterUser(name, email, password string) (*models.User, error) {
	// Hash and salt the password before storing it in the database
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create a new user with the hashed password
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	createdUser, err := um.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

// LoginUser handles the user login process
func (um *UserManager) LoginUser(email, password string) (*models.User, error) {
	// Implement your business logic for user login here
	// For simplicity, let's assume we're fetching the user from the database
	user, err := um.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Validate the password using bcrypt.CompareHashAndPassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// hashPassword hashes and salts a password
func hashPassword(password string) (string, error) {
	// Generate a hashed password with a cost factor of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
