package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

// SaveUserProfile saves a user's profile (avatar, bio)
func SaveUserProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.UserProfile
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SaveUserProfile(profile); err != nil {
		http.Error(w, "Error saving profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUserProfile retrieves a user's profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	profile, err := database.GetUserProfile(userID)
	if err != nil {
		http.Error(w, "Error fetching profile", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profile)
}
