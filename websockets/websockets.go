package websockets

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// WebSocket Upgrader: upgrades the HTTP connection to a WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (optional: implement stricter checks)
	},
}

// Struct for managing chat room clients
type Room struct {
	ID      string
	Clients map[*websocket.Conn]bool
	Mux     sync.Mutex
}

var rooms = make(map[string]*Room)

// JoinRoom handles WebSocket connections for chat rooms
func JoinRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_id")

	// Upgrade the connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}

	// Find or create the room
	room, exists := rooms[roomID]
	if !exists {
		room = &Room{
			ID:      roomID,
			Clients: make(map[*websocket.Conn]bool),
		}
		rooms[roomID] = room
	}

	// Add the client to the room
	room.Mux.Lock()
	room.Clients[conn] = true
	room.Mux.Unlock()

	log.Printf("User joined room: %s", roomID)

	// Listen for incoming messages
	go handleMessages(conn, room)
}

func handleMessages(conn *websocket.Conn, room *Room) {
	defer func() {
		// Cleanup when a user disconnects
		room.Mux.Lock()
		delete(room.Clients, conn)
		room.Mux.Unlock()
		conn.Close()
	}()

	for {
		// Read message from WebSocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Broadcast message to all clients in the room
		broadcastMessage(room, msg)
	}
}

// Broadcast message to all clients in the room
func broadcastMessage(room *Room, msg []byte) {
	room.Mux.Lock()
	defer room.Mux.Unlock()

	for conn := range room.Clients {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error sending message:", err)
			conn.Close()
			delete(room.Clients, conn)
		}
	}
}
