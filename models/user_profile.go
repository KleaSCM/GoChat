package models

type UserProfile struct {
	UserID string `json:"user_id"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}
