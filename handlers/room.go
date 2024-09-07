package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-SCM/gochat/database"
	"github.com/Jay-SCM/gochat/models"
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	roomID, err := database.CreateRoom(room)
	if err != nil {
		http.Error(w, "Failed to create room", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"room_id": roomID})
}

func GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := database.GetRooms()
	if err != nil {
		http.Error(w, "Failed to get rooms", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rooms)
}
