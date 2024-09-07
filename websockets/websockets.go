package websockets

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/yourusername/gochat/database"
	"github.com/yourusername/gochat/models"
)

type IncomingMessage struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
}

func broadcastMessage(room *Room, msg []byte) {
	room.Mux.Lock()
	defer room.Mux.Unlock()

	var incomingMsg IncomingMessage
	err := json.Unmarshal(msg, &incomingMsg)
	if err != nil {
		log.Println("Error parsing incoming message:", err)
		return
	}

	// Broadcast message to all clients in the room
	for conn := range room.Clients {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error sending message:", err)
			conn.Close()
			delete(room.Clients, conn)
		}
	}

	// Save the message to the database
	message := models.Message{
		RoomID:   incomingMsg.RoomID,
		Username: incomingMsg.Username,
		Content:  incomingMsg.Content,
	}
	if err := database.SaveMessage(message); err != nil {
		log.Println("Error saving message:", err)
	}
}
