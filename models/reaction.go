package models

type Reaction struct {
	MessageID string `json:"message_id"`
	UserID    string `json:"user_id"`
	Type      string `json:"type"`
}
