// repository/database.go

package repository

import (
   "fmt"
   "log"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/postgres"
   "github.com/kvn-media/SyncronizeHub/models"
   "github.com/kvn-media/SyncronizeHub/configs"  // Adjust the import path based on your project structure
)

// DB is the global database instance
var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() {
   cfg := config.NewConfig()

   // Build the database connection string
   dbURL := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s",
      cfg.DBUsername, cfg.DBName, cfg.DBPassword, cfg.DBSSLMode)

   var err error
   DB, err = gorm.Open("postgres", dbURL)

   if err != nil {
      log.Fatal(err)
   }

   // Set connection pool options if needed
   DB.DB().SetMaxIdleConns(10)
   DB.DB().SetMaxOpenConns(100)

   // AutoMigrate to create tables
   autoMigrateTables()
}

// AutoMigrateTables automigrates the tables
func autoMigrateTables() {
   // Migrate models to the database
   DB.AutoMigrate(&models.User{}, &models.FlowData{})
}

// CloseDB closes the database connection
func CloseDB() {
   if err := DB.Close(); err != nil {
      log.Fatal(err)
   }
}

// Other database-related functions, e.g., CRUD operations
// ...
