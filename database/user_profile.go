package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func SaveUserProfile(profile models.UserProfile) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET avatar = ?, bio = ? WHERE id = ?", profile.Avatar, profile.Bio, profile.UserID)
	if err != nil {
		log.Println("Failed to save user profile:", err)
		return err
	}
	return nil
}

func GetUserProfile(userID string) (models.UserProfile, error) {
	db, err := getDB()
	if err != nil {
		return models.UserProfile{}, err
	}

	var profile models.UserProfile
	err = db.QueryRow("SELECT avatar, bio FROM users WHERE id = ?", userID).Scan(&profile.Avatar, &profile.Bio)
	if err != nil {
		log.Println("Failed to get user profile:", err)
		return models.UserProfile{}, err
	}
	profile.UserID = userID
	return profile, nil
}
