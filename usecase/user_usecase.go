// usecase/user_usecase.go

package usecase

import (
	"fmt"
	"time"

	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/repository"
)

// UserUsecase handles business logic related to user operations.
type UserUsecase struct {
	userRepository *repository.UserRepository
}

// NewUserUsecase creates a new instance of UserUsecase.
func NewUserUsecase(userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

// RegisterUser registers a new user with the provided information.
func (uu *UserUsecase) RegisterUser(name, email, password string) (int, error) {
	// Additional validation logic can be added here before creating the user.
	// For example, check if the email is unique or if the password meets requirements.

	hashedPassword, err := hashAndSaltPassword(password)
	if err != nil {
		return 0, fmt.Errorf("error hashing and salting password")
	}

	user := models.NewUserModel(name, email, hashedPassword)

	userID, err := uu.userRepository.CreateUser(user)
	if err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}

	return userID, nil
}

// GetUserByID retrieves a user by ID.
func (uu *UserUsecase) GetUserByID(userID int) (*models.UserModel, error) {
	user, err := uu.userRepository.GetUserByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error getting user by ID: %v", err)
	}

	return user, nil
}

// UpdateUser updates user information.
func (uu *UserUsecase) UpdateUser(userID int, name, email string) error {
	// Additional validation logic can be added here before updating the user.
	// For example, check if the email is unique or if the user exists.

	user, err := uu.userRepository.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("error getting user by ID: %v", err)
	}

	user.Name = name
	user.Email = email
	user.UpdatedAt = time.Now()

	err = uu.userRepository.UpdateUser(user)
	if err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}

	return nil
}

// DeleteUser deletes a user by ID.
func (uu *UserUsecase) DeleteUser(userID int) error {
	// Additional logic can be added here before deleting the user.

	err := uu.userRepository.DeleteUser(userID)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}

// hashAndSaltPassword is a placeholder function for hashing and salting passwords.
func hashAndSaltPassword(password string) (string, error) {
	// Implement your password hashing and salting logic here.
	// This is a placeholder and should be replaced with a secure implementation.
	return password, nil
}
