// usecase/user_usecase.go

package usecase

import (
	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/managers"
)

// UserUseCase provides use cases for user-related operations
type UserUseCase struct {
	UserManager managers.UserManager
}

// NewUserUseCase creates a new instance of UserUseCase
func NewUserUseCase(userManager managers.UserManager) *UserUseCase {
	return &UserUseCase{
		UserManager: userManager,
	}
}

// RegisterUser handles the user registration use case
func (uc *UserUseCase) RegisterUser(name, email, password string) (*models.User, error) {
	return uc.UserManager.RegisterUser(name, email, password)
}

// LoginUser handles the user login use case
func (uc *UserUseCase) LoginUser(email, password string) (*models.User, error) {
	return uc.UserManager.LoginUser(email, password)
}
