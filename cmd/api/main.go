package main

import (
	"fmt"
	"net/http"

	// Import the new packages
	"homeplants-api/internal/database"
	"homeplants-api/internal/handlers"
)

func main() {
    // 1. Initialize DB (using the new package)
    database.InitDB()
	database.RunMigrations()
    defer database.DB.Close()

    // 2. Setup Routes (Using the handler package)
    http.HandleFunc("/rooms", handlers.GetRooms)

    http.HandleFunc("GET /rooms/{id}", handlers.GetRoomByID)
    http.HandleFunc("/plants", handlers.GetPlants)


    // 3. Start Server
    fmt.Println("Server starting on :8080 ...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}