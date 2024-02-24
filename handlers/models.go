 // why are byte arrays sometimes used instead of strings in go


package handlers

import (
	"golang.org/x/crypto/bcrypt"
)



// User represents a user's credentials

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}



// HashPassword takes a plain text password and returns a bcrypt hashed version
func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}


// CheckPasswordHash compares a bcrypt hashed password with its possible plaintext equivalent

func CheckPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}


