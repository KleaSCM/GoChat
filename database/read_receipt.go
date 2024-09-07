package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func SaveReadReceipt(receipt models.ReadReceipt) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO read_receipts (message_id, user_id, read_at) VALUES (?, ?, ?)", receipt.MessageID, receipt.UserID, receipt.ReadAt)
	if err != nil {
		log.Println("Failed to save read receipt:", err)
		return err
	}
	return nil
}

func GetReadReceipts(messageID string) ([]models.ReadReceipt, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT message_id, user_id, read_at FROM read_receipts WHERE message_id = ?", messageID)
	if err != nil {
		log.Println("Failed to get read receipts:", err)
		return nil, err
	}
	defer rows.Close()

	var receipts []models.ReadReceipt
	for rows.Next() {
		var receipt models.ReadReceipt
		if err := rows.Scan(&receipt.MessageID, &receipt.UserID, &receipt.ReadAt); err != nil {
			log.Println("Failed to scan read receipt:", err)
			return nil, err
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}
