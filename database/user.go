package database

import (
	"github.com/yourusername/gochat/models"
)

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := `SELECT id, username, password FROM users WHERE username = ?`
	err := DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
