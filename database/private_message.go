package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func GetPrivateMessages(userID string) ([]models.PrivateMessage, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, content, sender_id, receiver_id, created_at FROM private_messages WHERE receiver_id = ?", userID)
	if err != nil {
		log.Println("Failed to get private messages:", err)
		return nil, err
	}
	defer rows.Close()

	var messages []models.PrivateMessage
	for rows.Next() {
		var msg models.PrivateMessage
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.SenderID, &msg.ReceiverID, &msg.CreatedAt); err != nil {
			log.Println("Failed to scan private message:", err)
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}
