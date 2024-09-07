package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	RoomID    string    `json:"room_id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
