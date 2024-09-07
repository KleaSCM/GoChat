package models

type Notification struct {
	UserID string `json:"user_id"`
	RoomID string `json:"room_id"`
	IsRead bool   `json:"is_read"`
}
