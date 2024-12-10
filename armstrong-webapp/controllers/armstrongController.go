package controllers

import (
	"armstrong-webapp/database"
	"armstrong-webapp/models"
	"encoding/json"
	       // For formatted output (optional, if needed)
	"io"         // For reading the raw body
	"log"        // For logging debug information
	"net/http"
	"strconv"
	"strings"
)


// VerifyArmstrongNumber checks if a number is an Armstrong number and saves it for a user
// Debug log to inspect the incoming request body
// Debug log to inspect the incoming request body
func VerifyArmstrongNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
		Number int  `json:"number"`
	}

	// Log the raw body for debugging
	body := new(strings.Builder)
	_, err := io.Copy(body, r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	log.Println("Raw Body:", body.String()) // Log raw request body for debugging

	// Decode the JSON
	if err := json.NewDecoder(strings.NewReader(body.String())).Decode(&input); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	log.Printf("Decoded Input: UserID=%d, Number=%d\n", input.UserID, input.Number)

}


// GetUserArmstrongNumbers retrieves Armstrong numbers for a specific user
func GetUserArmstrongNumbers(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract user ID from the URL
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/user/"), "/numbers")
	if len(pathParts) == 0 || pathParts[0] == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(pathParts[0])
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch Armstrong numbers from the database
	var numbers []models.ArmstrongNumber
	result := database.DB.Where("user_id = ?", userID).Find(&numbers)
	if result.Error != nil {
		http.Error(w, "Database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
	log.Println("Numbers fetched from database:", numbers)

}

// GetAllUsersAndNumbers retrieves all users and their Armstrong numbers
func GetAllUsersAndNumbers(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fetch data with a JOIN query
	var results []struct {
		Email  string `json:"email"`
		Number int    `json:"number"`
	}
	database.DB.Raw(`
		SELECT users.email, armstrong_numbers.number
		FROM users
		JOIN armstrong_numbers ON users.id = armstrong_numbers.user_id
	`).Scan(&results)

	// Respond with the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
