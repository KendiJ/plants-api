package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "homeplants-api/internal/models"
)

// PlantHandler holds the dependencies. 
// No more global 'database.DB'! We inject it.
type PlantHandler struct {
    DB *sql.DB
}

func NewPlantHandler(db *sql.DB) *PlantHandler {
    return &PlantHandler{DB: db}
}

// Define the response structure as requested
type CreatePlantResponse struct {
    Message string `json:"message"`
    ID      int64  `json:"id"`
}

func (h *PlantHandler) CreatePlant(w http.ResponseWriter, r *http.Request) {
    var p models.Plant
    
    // Check for decoding errors
    if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    // Use Context for timeout (Addressing feedback)
    ctx := r.Context()

    query := `INSERT INTO plants (name, room_id, water_freq, image_url) VALUES (?, ?, ?, ?)`
    
    // Use h.DB (the injected database), not a global variable
    result, err := h.DB.ExecContext(ctx, query, p.Name, p.RoomID, p.WaterFreq, p.ImageURL)
    if err != nil {
        // Log the actual error internally, show generic to user
        // logger.Error("failed to insert plant", "error", err) 
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

    w.Header().Set("Content-Type", "application/json") // Fixed typo
    w.WriteHeader(http.StatusCreated)                  // Explicit status
    
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}