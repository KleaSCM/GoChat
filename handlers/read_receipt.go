package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

// SaveReadReceipt records a read receipt for a message
func SaveReadReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.ReadReceipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.SaveReadReceipt(receipt); err != nil {
		http.Error(w, "Error saving read receipt", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(receipt)
}

// GetReadReceipts retrieves all read receipts for a message
func GetReadReceipts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, err := strconv.Atoi(vars["message_id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	receipts, err := database.GetReadReceiptsByMessageID(messageID)
	if err != nil {
		http.Error(w, "Error fetching read receipts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(receipts)
}
