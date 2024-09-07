package database

import (
	"github.com/yourusername/gochat/models"
)

func CreateRoom(room models.Room) error {
	query := `INSERT INTO rooms (id, name) VALUES (?, ?)`
	_, err := DB.Exec(query, room.ID, room.Name)
	return err
}

func GetRooms() ([]models.Room, error) {
	query := `SELECT id, name FROM rooms`
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
