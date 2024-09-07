package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
)

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	messages, err := database.GetMessageHistory(roomID)
	if err != nil {
		http.Error(w, "Error fetching message history", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
