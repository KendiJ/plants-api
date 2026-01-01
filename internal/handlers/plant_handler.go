package handlers

import (
	"encoding/json"
	"homeplants-api/internal/database"
	"homeplants-api/internal/models"
	"net/http"
)
func GetPlants(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, room_id, water_freq FROM plants")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// emply slise for now
	plants := []models.Plant{}

	for rows.Next() {
		var p models.Plant
		if err := rows.Scan(&p.ID, &p.Name, &p.RoomID, &p.WaterFreq); err != nil {
			http.Error(w, "Error scanning plants", http.StatusInternalServerError)
			return
		}
		plants = append(plants, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plants)
}