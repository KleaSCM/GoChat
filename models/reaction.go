package models

import "time"

type Reaction struct {
	ID        int       `json:"id"`
	MessageID int       `json:"message_id"`
	Username  string    `json:"username"`
	Emoji     string    `json:"emoji"`
	CreatedAt time.Time `json:"created_at"`
}
