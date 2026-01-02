package handlers

import (
	"database/sql"
	"encoding/json"
	"homeplants-api/internal/database"
	"homeplants-api/internal/models"
	"net/http"
	"strconv"
)

// Capital 'G' in GetRooms exports it so main.go can use it
func GetRooms(w http.ResponseWriter, r *http.Request) {

	// 1. Get the ID from the URL
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		 return
	}

	// 2. query the DB (define the model locally)
	var room models.Room
	query := "SELECT id, name FROM rooms WHERE id = ?"
	err = database.DB.QueryRow(query, id).Scan(&room.ID, &room.Name)

	// 3. Handle "Not Found" vs "Real Error"
	if err == sql.ErrNoRows {
		http.Error(w, "Rooms not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	// 4 return JSON
	w.Header().Set("Content-type", "application.json")
	json.NewEncoder(w).Encode(room)


	//Use database.DB (The public variable we just made)
	rows, err := database.DB.Query("SELECT id, name FROM rooms")
	if err != nil {
		http.Error(w, "Database Error:", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Parse the rows into a slice
	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			http.Error(w, "Error Scanning room", http.StatusInternalServerError)
			return
		}
		rooms = append(rooms, room)
	}

	// Send JSON
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