package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func GetMessages(roomID string) ([]models.Message, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, content, sender_id, created_at FROM messages WHERE room_id = ?", roomID)
	if err != nil {
		log.Println("Failed to get messages:", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.CreatedAt); err != nil {
			log.Println("Failed to scan message:", err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
