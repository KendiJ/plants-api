package handlers

import (
	"database/sql"
	"encoding/json"
	"homeplants-api/internal/database"
	"homeplants-api/internal/models"
	"net/http"
	"strconv"
)


func GetRooms(w http.ResponseWriter, r *http.Request) {
	
	rows, err := database.DB.Query("SELECT id, name FROM rooms")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			continue
		}
		rooms = append(rooms, room)
	}

	
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(rooms)
}

func GetRoomByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
	}
	var room models.Room
	query := "SELECT id, name FROM rooms WHERE id = ?"

	err = database.DB.QueryRow(query, id).Scan(&room.ID, &room.Name)

	if err == sql.ErrNoRows {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(room)
}