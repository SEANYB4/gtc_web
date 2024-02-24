// File: database/db.go


package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)




// DB is the global database connection pool

var DB *sql.DB

// InitDB initializes the database connection using the given connection string.

func InitDB(dataSourceName string) {


	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}


	// Check the connection before proceeding
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %q", err)
	}

	fmt.Println("Connected to the database successfully!")
}


// CloseDB closes the database connection

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}