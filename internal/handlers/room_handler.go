package handlers

import (
	"encoding/json"
	"fmt"
	"homeplants-api/internal/database"
	"homeplants-api/internal/models"
	"net/http"
)

// Capital 'G' in GetRooms exports it so main.go can use it
func GetRooms(w http.ResponseWriter, r *http.Request) {
	//Use database.DB (The public variable we just made)
	rows, err := database.DB.Query("SELECT id, name FROM rooms")
	if err != nil {
		fmt.Println("❌ DB Query Error:", err)
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
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(rooms)
}