package handlers

import (
    "database/sql"
    "encoding/json"
    "homeplants-api/internal/models"
    "net/http"
    "strconv"
)

type RoomHandler struct {
    DB *sql.DB
}

func NewRoomHandler(db *sql.DB) *RoomHandler {
    return &RoomHandler{DB: db}
}

func (h *RoomHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
    rows, err := h.DB.Query("SELECT id, name, image_url FROM rooms")
    if err != nil {
        http.Error(w, "Database Error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    rooms := []models.Room{}
    
    for rows.Next() {
        var room models.Room
        if err := rows.Scan(&room.ID, &room.Name, &room.ImageURL); err != nil {
            continue
        }
        rooms = append(rooms, room)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(rooms)
}

func (h *RoomHandler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var room models.Room
    query := "SELECT id, name, image_url FROM rooms WHERE id = ?"

    err = h.DB.QueryRow(query, id).Scan(&room.ID, &room.Name, &room.ImageURL)

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

func (h *RoomHandler) GetPlantsByRoom(w http.ResponseWriter, r *http.Request) {
    idStr := r.PathValue("id")
    roomID, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid Room ID", http.StatusBadRequest)
        return
    }

    
    query := "SELECT id, name, room_id, water_freq, image_url FROM plants WHERE room_id = ?"
    rows, err := h.DB.Query(query, roomID)
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