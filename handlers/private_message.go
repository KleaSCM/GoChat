package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

// SendPrivateMessage handles sending private messages between users
func SendPrivateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.PrivateMessage
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SavePrivateMessage(msg); err != nil {
		http.Error(w, "Error saving message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

// GetPrivateMessages fetches message history between two users
func GetPrivateMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sender := vars["sender"]
	receiver := vars["receiver"]

	messages, err := database.GetPrivateMessages(sender, receiver)
	if err != nil {
		http.Error(w, "Error fetching messages", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(messages)
}
