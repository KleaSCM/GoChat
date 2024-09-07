package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func RegisterUser(user models.User) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Println("Failed to register user:", err)
		return err
	}
	return nil
}

func GetUser(username string) (models.User, error) {
	db, err := getDB()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println("Failed to get user:", err)
		return models.User{}, err
	}
	return user, nil
}
