package websockets

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type TypingIndicator struct {
	Username string `json:"username"`
	RoomID   string `json:"room_id"`
	Typing   bool   `json:"typing"`
}

func broadcastTypingStatus(room *Room, typingStatus TypingIndicator) {
	room.Mux.Lock()
	defer room.Mux.Unlock()

	statusMessage, err := json.Marshal(typingStatus)
	if err != nil {
		log.Println("Error marshalling typing status:", err)
		return
	}

	for conn := range room.Clients {
		err := conn.WriteMessage(websocket.TextMessage, statusMessage)
		if err != nil {
			log.Println("Error sending typing status:", err)
			conn.Close()
			delete(room.Clients, conn)
		}
	}
}

func handleMessages(conn *websocket.Conn, room *Room) {
	defer func() {
		room.Mux.Lock()
		delete(room.Clients, conn)
		room.Mux.Unlock()
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Check if it's a typing indicator
		var typingIndicator TypingIndicator
		if err := json.Unmarshal(msg, &typingIndicator); err == nil && typingIndicator.Typing {
			// Broadcast typing status
			broadcastTypingStatus(room, typingIndicator)
			continue
		}

		// Handle regular chat message
		broadcastMessage(room, msg)
	}
}
