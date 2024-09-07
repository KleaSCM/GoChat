package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Jay-SCM/gochat/database"
	"github.com/Jay-SCM/gochat/handlers"
	"github.com/Jay-SCM/gochat/websockets"
	"github.com/gorilla/mux"
)

func main() {
	// Start the background worker for archiving old messages
	go startBackgroundWorker()

	// Initialize the router
	router := mux.NewRouter()

	// Ensure the uploads directory exists
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		log.Fatalf("Failed to create uploads directory: %v", err)
	}

	// Set up routes
	setupRoutes(router)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRoutes(router *mux.Router) {
	// Public routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Chat room and message routes
	router.HandleFunc("/rooms", handlers.GetRooms).Methods("GET")
	router.HandleFunc("/rooms", handlers.CreateRoom).Methods("POST")
	router.HandleFunc("/rooms/{room_id}/messages", handlers.GetMessageHistory).Methods("GET")

	// Reactions and read receipts
	router.HandleFunc("/messages/{message_id}/reactions", handlers.GetReactions).Methods("GET")
	router.HandleFunc("/reactions", handlers.AddReaction).Methods("POST")
	router.HandleFunc("/messages/{message_id}/read_receipts", handlers.GetReadReceipts).Methods("GET")
	router.HandleFunc("/read_receipts", handlers.SaveReadReceipt).Methods("POST")

	// WebSocket route
	router.HandleFunc("/ws", websockets.JoinRoom).Methods("GET")

	// File upload route
	router.HandleFunc("/upload", handlers.UploadFile).Methods("POST")

	// Static file server for serving uploaded files
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	// Notifications routes
	router.HandleFunc("/notifications/{user_id}/{room_id}/unread", handlers.GetUnreadCount).Methods("GET")
	router.HandleFunc("/notifications/{user_id}/{room_id}/read", handlers.MarkMessagesAsRead).Methods("POST")

	// User profile routes
	router.HandleFunc("/profile", handlers.SaveUserProfile).Methods("POST")
	router.HandleFunc("/profile", handlers.GetUserProfile).Methods("GET")
}

func startBackgroundWorker() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop() // Ensure the ticker is stopped when the function exits

	for {
		select {
		case <-ticker.C:
			log.Println("Archiving old messages...")
			if err := database.ArchiveOldMessages(30); err != nil {
				log.Println("Error archiving messages:", err)
			} else {
				log.Println("Old messages archived successfully.")
			}
		}
	}
}
