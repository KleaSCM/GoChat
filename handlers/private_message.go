package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-SCM/gochat/database"
	"github.com/Jay-SCM/gochat/models"
)

func SendPrivateMessage(w http.ResponseWriter, r *http.Request) {
	var msg models.PrivateMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SendPrivateMessage(msg); err != nil {
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
