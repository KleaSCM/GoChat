package models

import "time"

type PrivateMessage struct {
	ID         string    `json:"id"`
	Content    string    `json:"content"`
	SenderID   string    `json:"sender_id"`
	ReceiverID string    `json:"receiver_id"`
	CreatedAt  time.Time `json:"created_at"`
}
