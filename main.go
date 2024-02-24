package main


import (
	"log"
	"net/http"
	"github.com/SEANYB4/gtc_web/handlers"
	"github.com/SEANYB4/gtc_web/database"
	"github.com/joho/godotenv"
	"os"
	"fmt"
)

func main() {

	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseHost := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")
	databaseUser := os.Getenv("DB_USER")
	databasePassword := os.Getenv("DB_PASSWORD")

	// Initialize the database connection
	// NOte: In production, it's better to use environment variables or a configuration file to store sensitive information such as db credentials
	dataSourceName := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=disable", databaseUser, databasePassword, databaseName, databaseHost)
	database.InitDB(dataSourceName)
	defer database.CloseDB()
	
	// Set up routes

	// Serve static files from the React app build directory
	fs := http.FileServer(http.Dir("../gtc_web_frontend/build"))
	http.Handle("/", fs)


	// user registration route
	http.HandleFunc("/register", handlers.RegisterUser)


	// user auth route
	http.HandleFunc("/login", handlers.Login)




	// Start the server

	log.Println("Starting server on:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}