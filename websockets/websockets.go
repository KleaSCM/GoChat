package websockets

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Presence struct {
	Username string `json:"username"`
	Online   bool   `json:"online"`
}

func broadcastPresence(room *Room, presence Presence) {
	room.Mux.Lock()
	defer room.Mux.Unlock()

	statusMessage, err := json.Marshal(presence)
	if err != nil {
		log.Println("Error marshalling presence status:", err)
		return
	}

	for conn := range room.Clients {
		err := conn.WriteMessage(websocket.TextMessage, statusMessage)
		if err != nil {
			log.Println("Error sending presence status:", err)
			conn.Close()
			delete(room.Clients, conn)
		}
	}
}

func handleConnection(conn *websocket.Conn, room *Room, username string) {
	defer func() {
		room.Mux.Lock()
		delete(room.Clients, conn)
		room.Mux.Unlock()

		// Broadcast user is offline
		broadcastPresence(room, Presence{Username: username, Online: false})
		conn.Close()
	}()

	// Broadcast user is online
	broadcastPresence(room, Presence{Username: username, Online: true})

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
	}
}
