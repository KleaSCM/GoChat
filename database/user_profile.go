package database

import "github.com/yourusername/gochat/models"

func SaveUserProfile(profile models.UserProfile) error {
	query := `INSERT INTO user_profiles (user_id, avatar, bio) VALUES (?, ?, ?)
              ON DUPLICATE KEY UPDATE avatar = ?, bio = ?`
	_, err := DB.Exec(query, profile.UserID, profile.Avatar, profile.Bio, profile.Avatar, profile.Bio)
	return err
}

func GetUserProfile(userID int) (models.UserProfile, error) {
	var profile models.UserProfile
	query := `SELECT user_id, avatar, bio FROM user_profiles WHERE user_id = ?`
	err := DB.QueryRow(query, userID).Scan(&profile.UserID, &profile.Avatar, &profile.Bio)
	if err != nil {
		return models.UserProfile{}, err
	}
	return profile, nil
}
