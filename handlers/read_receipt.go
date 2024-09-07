package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-SCM/gochat/database"
	"github.com/Jay-SCM/gochat/models"
)

func SaveReadReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.ReadReceipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SaveReadReceipt(receipt); err != nil {
		http.Error(w, "Failed to save read receipt", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
