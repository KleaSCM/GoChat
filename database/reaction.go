package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func AddReaction(reaction models.Reaction) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO reactions (message_id, user_id, type) VALUES (?, ?, ?)", reaction.MessageID, reaction.UserID, reaction.Type)
	if err != nil {
		log.Println("Failed to add reaction:", err)
		return err
	}
	return nil
}

func GetReactions(messageID string) ([]models.Reaction, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT message_id, user_id, type FROM reactions WHERE message_id = ?", messageID)
	if err != nil {
		log.Println("Failed to get reactions:", err)
		return nil, err
	}
	defer rows.Close()

	var reactions []models.Reaction
	for rows.Next() {
		var reaction models.Reaction
		if err := rows.Scan(&reaction.MessageID, &reaction.UserID, &reaction.Type); err != nil {
			log.Println("Failed to scan reaction:", err)
			return nil, err
		}
		reactions = append(reactions, reaction)
	}
	return reactions, nil
}
