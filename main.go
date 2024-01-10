// main.go

package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kvn-media/SyncronizeHub/delivery/controllers"
	"github.com/kvn-media/SyncronizeHub/managers"
	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/repository"
)

func main() {
	// Initialize Gorm database connection (replace with your actual database connection details)
	db, err := gorm.Open("postgres", "user=postgres dbname=SyncronizeHub sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// AutoMigrate to create tables
	db.AutoMigrate(&models.User{}, &models.FlowData{})

	// Initialize UserRepository
	userRepository := repository.NewUserRepository(db)

	// Initialize UserManager and UserController
	userManager := managers.NewUserManager(userRepository)
	userController := controllers.NewUserController(userManager)

	// Create a new Gorilla Mux router
	r := mux.NewRouter()

	// Define routes and associate them with UserController methods
	r.HandleFunc("/public", userController.PublicEndpoint).Methods("GET")
	r.HandleFunc("/register", userController.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userController.LoginUser).Methods("POST")

	// Other routes and handlers can be added here

	// Use the router as the main handler
	http.Handle("/", r)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
