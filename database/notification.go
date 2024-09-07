package database

func IncrementUnreadCount(userID, roomID int) error {
	query := `INSERT INTO unread_notifications (user_id, room_id, unread_count)
              VALUES (?, ?, 1) ON DUPLICATE KEY UPDATE unread_count = unread_count + 1`
	_, err := DB.Exec(query, userID, roomID)
	return err
}

func GetUnreadCount(userID, roomID int) (int, error) {
	var unreadCount int
	query := `SELECT unread_count FROM unread_notifications WHERE user_id = ? AND room_id = ?`
	err := DB.QueryRow(query, userID, roomID).Scan(&unreadCount)
	if err != nil {
		return 0, err
	}
	return unreadCount, nil
}

func MarkMessagesAsRead(userID, roomID int) error {
	query := `UPDATE unread_notifications SET unread_count = 0 WHERE user_id = ? AND roomID = ?`
	_, err := DB.Exec(query, userID, roomID)
	return err
}
