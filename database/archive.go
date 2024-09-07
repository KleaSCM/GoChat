package database

import (
	"time"
)

func ArchiveOldMessages(days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)
	query := `DELETE FROM messages WHERE created_at < ?`
	_, err := DB.Exec(query, cutoff)
	return err
}
