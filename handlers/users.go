package handlers



import (
	"encoding/json"
	"net/http"
	// Assume "github.com/SEANYB4/gtc_web/database" is your package that handles DB operations
	"github.com/SEANYB4/gtc_web/database"
)




// RegisterUser handles the user registration process
func RegisterUser(w http.ResponseWriter, r *http.Request) {


	if r.Method != "POST" {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}


	var user User
	err := json.NewDecoder(r.Body).Decode(&user)


	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	// Hash the user's password
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Create a new User object for storing to the DB

	dbUser := database.User{
		Username: user.Username,
		Password: hashedPassword,
	}

	// Insert the user into the database
	err = database.InsertUser(dbUser)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}


	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))

	
}