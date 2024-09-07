package models

type UnreadNotification struct {
	UserID      int `json:"user_id"`
	RoomID      int `json:"room_id"`
	UnreadCount int `json:"unread_count"`
}
