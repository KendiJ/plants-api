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

func CreatePlant(w http.ResponseWriter, r *http.Request) {
	// 1. Define the expected structure (What Swift sends us)
	var p struct {
		Name string `json:"name"`
		RoomID string `json:"room_id"`
		WaterFreq string `json:"water_frequency"`
		ImageURL string `json:"image_url"`
	}

	// 2. Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//3. Validation (check if name is too short or empty)
	if len(p.Name) < 2 {
		http.Error(w, "Plant name is too short", http.StatusBadRequest)
		return
	}

	// 4. Insert into Database
	query := `INSERT INTO plants(name, room_id, water_freq, image_url) VALUES (?, ?, ?, ?)`
	result, err := database.DB.Exec(query, p.Name, p.RoomID, p.WaterFreq, p.ImageURL)
	if err != nil {
		http.Error(w, "Failed to save plant", http.StatusInternalServerError)
		return
	}

	// 5. Send back the new ID 
	id, _ := result.LastInsertId()
	w.Header().Set("Content type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Plant created successfully",
		"id": id,
	})
}