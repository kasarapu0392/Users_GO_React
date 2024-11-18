package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
	// Database connection parameters
	dbUser := "myuser"
	dbPassword := "mypassword"
	dbName := "myapp"
	dbHost := "localhost" // Default host
	dbPort := "5432"      // Default PostgreSQL port

	// Construct connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a connection to the database
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("[InitDB] Failed to connect to database: %v", err)
	}

	// Ping the database to verify connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("[InitDB] Database ping failed: %v", err)
	}

	log.Println("[InitDB] Database connection established successfully")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("[CloseDB] Error closing database connection: %v", err)
		} else {
			log.Println("[CloseDB] Database connection closed successfully")
		}
	}
}
