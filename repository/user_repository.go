// repository/user_repository.go

package repository

import "github.com/kvn-media/SyncronizeHub/models"

// UserRepository defines methods to interact with the user data store
type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	// Add other user-related methods as needed
}
