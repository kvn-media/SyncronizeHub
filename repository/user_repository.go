// repository/user_repository.go

package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kvn-media/SyncronizeHub/models"
)

// UserRepository handles data access operations for the UserModel.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// CreateUser inserts a new user into the database.
func (ur *UserRepository) CreateUser(user *models.UserModel) (int, error) {
	query := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	var userID int
	err := ur.DB.QueryRow(query, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt).Scan(&userID)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return 0, fmt.Errorf("error creating user")
	}

	return userID, nil
}

// GetUserByID retrieves a user from the database by ID.
func (ur *UserRepository) GetUserByID(userID int) (*models.UserModel, error) {
	query := `
		SELECT id, name, email, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	var user models.UserModel
	err := ur.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error getting user by ID: %v\n", err)
		return nil, fmt.Errorf("error getting user by ID")
	}

	return &user, nil
}

// UpdateUser updates a user in the database.
func (ur *UserRepository) UpdateUser(user *models.UserModel) error {
	query := `
		UPDATE users
		SET name = $2, email = $3, updated_at = $4
		WHERE id = $1
	`

	_, err := ur.DB.Exec(query, user.ID, user.Name, user.Email, user.UpdatedAt)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
		return fmt.Errorf("error updating user")
	}

	return nil
}

// DeleteUser deletes a user from the database.
func (ur *UserRepository) DeleteUser(userID int) error {
	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := ur.DB.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting user: %v\n", err)
		return fmt.Errorf("error deleting user")
	}

	return nil
}
