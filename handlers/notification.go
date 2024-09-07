package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/database"
)

// GetUnreadCount returns the unread message count for a user in a chat room
func GetUnreadCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	roomID, _ := strconv.Atoi(vars["room_id"])

	count, err := database.GetUnreadCount(userID, roomID)
	if err != nil {
		http.Error(w, "Unable to get unread count", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"unread_count": count})
}

// MarkMessagesAsRead marks all messages as read for a user in a chat room
func MarkMessagesAsRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["user_id"])
	roomID, _ := strconv.Atoi(vars["room_id"])

	if err := database.MarkMessagesAsRead(userID, roomID); err != nil {
		http.Error(w, "Unable to mark messages as read", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
