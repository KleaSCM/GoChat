package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

// AddReaction allows users to react to a message
func AddReaction(w http.ResponseWriter, r *http.Request) {
	var reaction models.Reaction
	err := json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.AddReaction(reaction); err != nil {
		http.Error(w, "Error adding reaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reaction)
}

// GetReactions retrieves reactions for a specific message
func GetReactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	messageID, err := strconv.Atoi(vars["message_id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	reactions, err := database.GetReactionsByMessageID(messageID)
	if err != nil {
		http.Error(w, "Error fetching reactions", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reactions)
}
