// delivery/controllers/app_controller.go

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/usecase"
)

// AppController handles HTTP requests related to application functionality.
type AppController struct {
	userUsecase *usecase.UserUsecase
}

// NewAppController creates a new instance of AppController.
func NewAppController(userUsecase *usecase.UserUsecase) *AppController {
	return &AppController{
		userUsecase: userUsecase,
	}
}

// RegisterUser handles the registration of a new user.
func (ac *AppController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var userRequest models.UserModel

	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		handleError(w, http.StatusBadRequest, "invalid request payload")
		return
	}

	userID, err := ac.userUsecase.RegisterUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("error registering user: %v", err))
		return
	}

	handleSuccess(w, http.StatusCreated, map[string]int{"user_id": userID})
}

// GetUserByID handles the retrieval of a user by ID.
func (ac *AppController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		handleError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	user, err := ac.userUsecase.GetUserByID(userID)
	if err != nil {
		handleError(w, http.StatusInternalServerError, fmt.Sprintf("error getting user by ID: %v", err))
		return
	}

	handleSuccess(w, http.StatusOK, user)
}

// handleSuccess responds with a success JSON response.
func handleSuccess(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// handleError responds with an error JSON response.
func handleError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
