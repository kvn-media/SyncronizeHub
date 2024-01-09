// managers/infra_manager.go

package managers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// InfraManager handles infrastructure-related operations, such as database connections.
type InfraManager struct {
	DB *sql.DB
}

// NewInfraManager creates a new instance of InfraManager.
func NewInfraManager(dbConfig string) (*InfraManager, error) {
	db, err := sql.Open("postgres", dbConfig)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Connected to the database")

	return &InfraManager{
		DB: db,
	}, nil
}

// CloseDB closes the database connection.
func (im *InfraManager) CloseDB() {
	if im.DB != nil {
		err := im.DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v\n", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}

// MigrateDB performs database migrations.
func (im *InfraManager) MigrateDB() error {
	// Example migration logic (replace with actual migrations)
	_, err := im.DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255) UNIQUE,
			password VARCHAR(255)
		);
		-- Add more migrations as needed
	`)
	if err != nil {
		return fmt.Errorf("error executing database migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
	return nil
}
