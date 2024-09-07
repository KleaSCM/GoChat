package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/gochat/handlers"
	"github.com/yourusername/gochat/middlewares"
)

func main() {
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Apply rate limiting middleware to message routes
	messageRouter := router.PathPrefix("/messages").Subrouter()
	messageRouter.Use(middlewares.RateLimit)
	messageRouter.HandleFunc("/{room_id}/messages", handlers.GetMessageHistory).Methods("GET")
	messageRouter.HandleFunc("/reactions", handlers.AddReaction).Methods("POST")

	// WebSocket route
	router.HandleFunc("/ws", websockets.JoinRoom).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
