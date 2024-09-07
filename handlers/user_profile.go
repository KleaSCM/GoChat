package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

func SaveUserProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.UserProfile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SaveUserProfile(profile); err != nil {
		http.Error(w, "Failed to save user profile", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	profile, err := database.GetUserProfile(userID)
	if err != nil {
		http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
