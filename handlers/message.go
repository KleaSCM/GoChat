package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-SCM/gochat/database"
	"github.com/gorilla/mux"
)

func GetMessageHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	roomID := vars["room_id"]

	messages, err := database.GetMessages(roomID)
	if err != nil {
		http.Error(w, "Failed to get messages", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}
