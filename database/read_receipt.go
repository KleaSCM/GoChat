package database

import (
	"time"

	"github.com/yourusername/gochat/models"
)

func SaveReadReceipt(receipt models.ReadReceipt) error {
	query := `INSERT INTO read_receipts (message_id, username, read_at) VALUES (?, ?, ?)`
	_, err := DB.Exec(query, receipt.MessageID, receipt.Username, time.Now())
	return err
}

func GetReadReceiptsByMessageID(messageID int) ([]models.ReadReceipt, error) {
	query := `SELECT message_id, username, read_at FROM read_receipts WHERE message_id = ?`
	rows, err := DB.Query(query, messageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var receipts []models.ReadReceipt
	for rows.Next() {
		var receipt models.ReadReceipt
		if err := rows.Scan(&receipt.MessageID, &receipt.Username, &receipt.ReadAt); err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}

	return receipts, nil
}
