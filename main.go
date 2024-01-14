package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/kvn-media/SyncronizeHub/configs"
	"github.com/kvn-media/SyncronizeHub/delivery/controllers"
	"github.com/kvn-media/SyncronizeHub/managers"
	"github.com/kvn-media/SyncronizeHub/repository"
)

func main() {
	// Initialize database connection and perform auto-migration
	err, _ := repository.InitDB(*config.NewConfig())
	if err != nil {
		panic(err)
	}
	defer repository.DB.Close()

	// Initialize UserRepository
	userRepository := repository.NewUserRepository(repository.DB)

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
