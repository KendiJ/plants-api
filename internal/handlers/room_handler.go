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
	rows, err := h.DB.Query("SELECT id, name, image_urls FROM rooms")
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Name, &room.ImageURL); err != nil {
			continue
		}
		rooms = append(rooms, room)
	}

	w.Header().Set("Content Type", "application/json")
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
	query := "SELECT id, name FROM rooms WHERE id = ?"

	err = h.DB.QueryRow(query, id).Scan(&room.ID, &room.Name)

	if err == sql.ErrNoRows {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(room)
}

func (h *RoomHandler) GetPlantsByRoom(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}