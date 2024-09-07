package database

import (
	"log"
)

func GetUnreadCount(userID, roomID string) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM notifications WHERE user_id = ? AND room_id = ? AND is_read = FALSE", userID, roomID).Scan(&count)
	if err != nil {
		log.Println("Failed to get unread count:", err)
		return 0, err
	}
	return count, nil
}
