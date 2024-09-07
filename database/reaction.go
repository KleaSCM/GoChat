package database

import (
	"time"

	"github.com/yourusername/gochat/models"
)

func AddReaction(reaction models.Reaction) error {
	query := `INSERT INTO reactions (message_id, username, emoji, created_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, reaction.MessageID, reaction.Username, reaction.Emoji, time.Now())
	return err
}

func GetReactionsByMessageID(messageID int) ([]models.Reaction, error) {
	query := `SELECT id, message_id, username, emoji, created_at FROM reactions WHERE message_id = ?`
	rows, err := DB.Query(query, messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reactions []models.Reaction
	for rows.Next() {
		var reaction models.Reaction
		if err := rows.Scan(&reaction.ID, &reaction.MessageID, &reaction.Username, &reaction.Emoji, &reaction.CreatedAt); err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	return reactions, nil
}
