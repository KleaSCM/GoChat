package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/handlers"
	"github.com/yourusername/gochat/websockets"
)

func main() {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Chat room routes
	router.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	router.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")

	// Authenticated WebSocket route
	router.HandleFunc("/ws", websockets.JoinRoom).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}