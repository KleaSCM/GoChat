package database

import (
	"database/sql"
	"log"
	"time"
)

func ArchiveOldMessages(days int) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	cutoffDate := time.Now().AddDate(0, 0, -days)
	_, err = db.Exec("DELETE FROM messages WHERE created_at < ?", cutoffDate)
	if err != nil {
		log.Println("Failed to archive old messages:", err)
		return err
	}
	return nil
}

func getDB() (*sql.DB, error) {
	// Database connection logic here
	return nil, nil
}
