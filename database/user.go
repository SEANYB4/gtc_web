package database

import (
	"log"
)


// User represents the user's credentials stored in the database
type User struct {
	Username string
	Password string
}



// INsertUser inserts a new user into the database

func InsertUser(u User) error {

	_, err := DB.Exec("INSERT INTO users(username, password) VALUES($1, $2)", u.Username, u.Password)
	if err != nil {
		log.Printf("Error inserting new user: %v", err)
		return err
	}
	return nil
}