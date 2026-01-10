package handlers

import (
	"database/sql"
	"encoding/json"
	"homeplants-api/internal/database"
	"homeplants-api/internal/models"
	"net/http"
	"strconv"
)
func GetPlants(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, room_id, water_freq, image_url FROM plants")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// emply slise for now
	plants := []models.Plant{}

	for rows.Next() {
		var p models.Plant
		if err := rows.Scan(&p.ID, &p.Name, &p.RoomID, &p.WaterFreq, &p.ImageURL); err != nil {
			http.Error(w, "Error scanning plants", http.StatusInternalServerError)
			return
		}
		plants = append(plants, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plants)
}

func GetPlantByID (w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p models.Plant
	query := "SELECT id, name, room_id, water_freq, image_url FROM plants WHERE id = ?"

	err = database.DB.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.RoomID, &p.WaterFreq, &p.ImageURL)

	if err == sql.ErrNoRows {
		http.Error(w, "Plant not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database erro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func GetPlantsByRoom(w http.ResponseWriter, r *http.Request) {
	roomIDStr := r.PathValue("id")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		http.Error(w, "Invalid Room ID", http.StatusBadRequest)
		return
	}

	
	query := "SELECT id, name, room_id, water_freq, image_url FROM plants WHERE room_id = ?"
	rows, err := database.DB.Query(query, roomID)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	plants := []models.Plant{}
	for rows.Next() {
		var p models.Plant
		if err := rows.Scan(&p.ID, &p.Name, &p.RoomID, &p.WaterFreq, &p.ImageURL); err != nil {
			continue
		}
		plants = append(plants, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plants)
}