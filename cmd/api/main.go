package main

import (
	"fmt"
	"net/http"

	
	"homeplants-api/internal/database"
	"homeplants-api/internal/handlers"
)

func main() {
    
    database.InitDB()
	database.RunMigrations()
    defer database.DB.Close()

   
    http.HandleFunc("GET /rooms", handlers.GetRooms)
    http.HandleFunc("GET /rooms/{id}", handlers.GetRoomByID)
    http.HandleFunc("GET /rooms/{id}/plants", handlers.GetPlantsByRoom)

    
    http.HandleFunc("GET /plants", handlers.GetPlants)
    http.HandleFunc("GET /plants/{id}", handlers.GetPlantByID)
    http.HandleFunc("POST /plants", handlers.CreatePlant)


    
    fmt.Println("Server starting on :8080 ...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}