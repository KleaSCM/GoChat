package main

import (
	"log"
	"time"

	"github.com/yourusername/gochat/database"
)

func startBackgroundWorker() {
	ticker := time.NewTicker(24 * time.Hour)
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
