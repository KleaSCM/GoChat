package models

import "time"

type ReadReceipt struct {
	MessageID int       `json:"message_id"`
	Username  string    `json:"username"`
	ReadAt    time.Time `json:"read_at"`
}
