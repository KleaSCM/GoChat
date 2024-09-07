package models

type UserProfile struct {
	UserID int    `json:"user_id"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
}
