// repository/user_repository.go

package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kvn-media/SyncronizeHub/models"
)

// UserRepository defines methods to interact with user data
type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// userDBRepository implements UserRepository using a Gorm database connection
type userDBRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of userDBRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userDBRepository{
		db: db,
	}
}

// CreateUser creates a new user in the database
func (ur *userDBRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail retrieves a user by email from the database
func (ur *userDBRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}