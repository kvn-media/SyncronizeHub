// main.go

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kvn-media/SyncronizeHub/controllers"
	"github.com/kvn-media/SyncronizeHub/managers"
)

func main() {
	// Initialize UserManager and UserController
	userManager := managers.NewUserManager()
	userController := controllers.NewUserController(userManager)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define routes and associate them with UserController methods
	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userController.LoginUser).Methods("POST")

	// Other routes and handlers can be added here

	// Use the router as the main handler
	http.Handle("/", r)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
