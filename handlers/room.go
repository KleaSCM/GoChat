package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

// CreateRoom handles the creation of new chat rooms
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	room.ID = uuid.New().String() // Generate a unique room ID
	err = database.CreateRoom(room)
	if err != nil {
		http.Error(w, "Error creating room", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(room)
}

// GetRooms returns a list of available chat rooms
func GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := database.GetRooms()
	if err != nil {
		http.Error(w, "Error fetching rooms", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rooms)
}
