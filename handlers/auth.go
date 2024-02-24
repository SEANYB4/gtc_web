package handlers


import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
)


// SimpleHashPassword hashes the password using SHA-256 (just for example purposes, use bcrypt in production)


func SimpleHashPassword(password string) string {

	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}



// SimpleTokenGeneration (just for example purposes, use JWT or similar in production)


func SimpleTokenGeneration(userID string) string {

	// This is a placeholder for a proper token generation

	hash := sha256.Sum256([]byte(userID))
	return hex.EncodeToString(hash[:])
}



func Login(w http.ResponseWriter, r *http.Request) {
	// Youir login logic goes here

	// Placeholder for token generation

	token := SimpleTokenGeneration("userID")



	w.Header().Set("Authorization", "Bearer "+token)
	w.Write([]byte("Logged in successfully"))


}