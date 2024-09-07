package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-SCM/gochat/database"
	"github.com/Jay-SCM/gochat/models"
)

func AddReaction(w http.ResponseWriter, r *http.Request) {
	var reaction models.Reaction
	if err := json.NewDecoder(r.Body).Decode(&reaction); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := database.AddReaction(reaction); err != nil {
		http.Error(w, "Failed to add reaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
