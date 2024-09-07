package database

import (
	"time"

	"github.com/yourusername/gochat/models"
)

func SavePrivateMessage(msg models.PrivateMessage) error {
	query := `INSERT INTO private_messages (sender, receiver, content, created_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, msg.Sender, msg.Receiver, msg.Content, time.Now())
	return err
}

func GetPrivateMessages(sender, receiver string) ([]models.PrivateMessage, error) {
	query := `SELECT id, sender, receiver, content, created_at FROM private_messages 
              WHERE (sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)
              ORDER BY created_at ASC`
	rows, err := DB.Query(query, sender, receiver, receiver, sender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.PrivateMessage
	for rows.Next() {
		var msg models.PrivateMessage
		if err := rows.Scan(&msg.ID, &msg.Sender, &msg.Receiver, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
