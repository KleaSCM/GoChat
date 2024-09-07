package models

import "time"

type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	SenderID  string    `json:"sender_id"`
	CreatedAt time.Time `json:"created_at"`
}
