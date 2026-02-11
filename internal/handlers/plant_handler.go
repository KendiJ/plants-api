package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "homeplants-api/internal/models"
)


type PlantHandler struct {
    DB *sql.DB
}

func NewPlantHandler(db *sql.DB) *PlantHandler {
    return &PlantHandler{DB: db}
}

type CreatePlantResponse struct {
    Message string `json:"message"`
    ID      int64  `json:"id"`
}

func (h *PlantHandler) CreatePlant(w http.ResponseWriter, r *http.Request) {
    var p models.Plant
    
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    ctx := r.Context()

    query := `INSERT INTO plants (name, room_id, water_freq, image_url) VALUES (?, ?, ?, ?)`
    
    result, err := h.DB.ExecContext(ctx, query, p.Name, p.RoomID, p.WaterFreq, p.ImageURL)
    if err != nil {
        http.Error(w, "Failed to create plant", http.StatusInternalServerError)
        return
    }

    id, err := result.LastInsertId()
    if err != nil {
        http.Error(w, "Failed to get new plant ID", http.StatusInternalServerError)
        return
    }

    response := CreatePlantResponse{
        Message: "Plant created successfully",
        ID:      id,
    }

    w.Header().Set("Content-Type", "application/json") 
    w.WriteHeader(http.StatusCreated)                  
    
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}