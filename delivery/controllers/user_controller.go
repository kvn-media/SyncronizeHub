// delivery/controllers/user_controller.go

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/kvn-media/SyncronizeHub/managers"
	"github.com/kvn-media/SyncronizeHub/models"
)

// UserController handles HTTP requests related to users
type UserController struct {
	UserManager *managers.UserManager
}

// NewUserController creates a new instance of UserController
func NewUserController(userManager *managers.UserManager) *UserController {
	return &UserController{
		UserManager: userManager,
	}
}

// RegisterUser handles user registration
func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call UserManager to register the user
	registeredUser, err := uc.UserManager.RegisterUser(newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	// Return the registered user as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(registeredUser)
}

// LoginUser handles user login
func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var loginRequest models.UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call UserManager to login the user
	loggedInUser, err := uc.UserManager.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return the logged-in user as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loggedInUser)
}

// Add more controller functions as needed for your project
func (uc *UserController) PublicEndpoint(w http.ResponseWriter, r *http.Request) {
	
}