package database

import (
	"time"

	"github.com/yourusername/gochat/models"
)

func SaveMessage(msg models.Message) error {
	query := `INSERT INTO messages (room_id, username, content, created_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, msg.RoomID, msg.Username, msg.Content, time.Now())
	return err
}

func GetMessageHistory(roomID string) ([]models.Message, error) {
	query := `SELECT id, room_id, username, content, created_at FROM messages WHERE room_id = ? ORDER BY created_at ASC`
	rows, err := DB.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.RoomID, &msg.Username, &msg.Content, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	return messages, nil
}
