package database

import (
	"log"

	"github.com/Jay-SCM/gochat/models"
)

func CreateRoom(room models.Room) (string, error) {
	db, err := getDB()
	if err != nil {
		return "", err
	}

	result, err := db.Exec("INSERT INTO rooms (name) VALUES (?)", room.Name)
	if err != nil {
		log.Println("Failed to create room:", err)
		return "", err
	}

	roomID, err := result.LastInsertId()
	if err != nil {
		log.Println("Failed to get room ID:", err)
		return "", err
	}
	return string(roomID), nil
}

func GetRooms() ([]models.Room, error) {
	db, err := getDB()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, name FROM rooms")
	if err != nil {
		log.Println("Failed to get rooms:", err)
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			log.Println("Failed to scan room:", err)
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
