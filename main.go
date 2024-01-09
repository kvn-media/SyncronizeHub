// main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/kvn-media/SyncronizeHub/configs"
	"github.com/kvn-media/SyncronizeHub/delivery"
	"github.com/kvn-media/SyncronizeHub/delivery/controllers"
	"github.com/kvn-media/SyncronizeHub/delivery/middleware"
	"github.com/kvn-media/SyncronizeHub/managers"
	"github.com/kvn-media/SyncronizeHub/models"
	"github.com/kvn-media/SyncronizeHub/repository"
	"github.com/kvn-media/SyncronizeHub/usecase"
	"github.com/kvn-media/SyncronizeHub/utils/authenticator"
	"github.com/kvn-media/SyncronizeHub/utils/common_responses"
	"github.com/kvn-media/SyncronizeHub/utils/generate_id"
	"github.com/gorilla/mux"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5432"
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_DRIVER   = "postgres"

	API_HOST = "localhost"
	API_PORT = "8989"

	APP_NAME = ""

	FirebaseCredentialsPath = "path/to/credentials.json"
	FirebaseDatabaseURL     = "https://your-project-id.firebaseio.com"
)

func main() {
	// Setup environment variables and view configurations
	setEnv()
	viewConfigs()

	// Initialize Firebase Manager
	firebaseManager, err := authenticator.NewFirebaseManager(FirebaseCredentialsPath, FirebaseDatabaseURL)
	if err != nil {
		log.Fatalf("Error initializing Firebase Manager: %v", err)
	}

	// Initialize Paseto Manager
	pasetoManager := authenticator.NewPasetoManager([]byte("your_secret_key"))

	// Initialize other components
	config := configs.NewConfig()
	generateID := generate_id.NewGenerateID()

	// Initialize database
	dbManager := managers.NewDatabaseManager(DB_DRIVER, getDBConnectionString())
	dbManager.SetupDatabase()

	// Initialize Repositories
	userRepo := repository.NewUserRepository(dbManager.DB)
	flowDataRepo := repository.NewFlowDataRepository(dbManager.DB)

	// Initialize Use Cases
	userUsecase := usecase.NewUserUsecase(userRepo, generateID, pasetoManager, config)
	flowDataUsecase := usecase.NewFlowDataUsecase(flowDataRepo, generateID, config)

	// Initialize Controllers
	userController := controllers.NewUserController(userUsecase)
	flowDataController := controllers.NewFlowDataController(flowDataUsecase)

	// Initialize Middleware
	authMiddleware := middleware.NewAuthMiddleware(pasetoManager)

	// Initialize Router
	router := mux.NewRouter()

	// Apply Middleware
	router.Use(authMiddleware.Middleware)

	// Define API Endpoints
	router.HandleFunc("/api/user/register", userController.Register).Methods("POST")
	router.HandleFunc("/api/user/login", userController.Login).Methods("POST")
	router.HandleFunc("/api/flowdata", flowDataController.GetFlowData).Methods("GET")

	// Serve the API
	apiAddr := fmt.Sprintf("%s:%s", API_HOST, API_PORT)
	apiServer := &http.Server{
		Addr:    apiAddr,
		Handler: router,
	}

	log.Printf("Starting %s API Server at %s\n", APP_NAME, apiAddr)
	log.Fatal(apiServer.ListenAndServe())
}

func setEnv() {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Auto setting environment variables")
	fmt.Println("You can disable this feature on main.go")

	os.Setenv("DB_HOST", DB_HOST)
	os.Setenv("DB_PORT", DB_PORT)
	os.Setenv("DB_USER", DB_USER)
	os.Setenv("DB_PASSWORD", DB_PASSWORD)
	os.Setenv("DB_NAME", DB_NAME)
	os.Setenv("DB_DRIVER", DB_DRIVER)

	os.Setenv("API_HOST", API_HOST)
	os.Setenv("API_PORT", API_PORT)

	os.Setenv("APP_NAME", APP_NAME)

	fmt.Println("Setting finished")
	fmt.Println(strings.Repeat("=", 50))
}

func viewConfigs() {
	config := configs.NewConfig()

	fmt.Println("configs: ")
	fmt.Println("db config:", config.DbConfig)
	fmt.Println("api config (port maybe auto set):", config.ApiConfig)

	fmt.Println(strings.Repeat("=", 50))

	fmt.Println()
}

func getDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
