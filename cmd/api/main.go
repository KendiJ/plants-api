package main

import (
	"fmt"
	"log"
	"net/http"
    

	// "homeplants-api/internal/config"
    "github.com/joho/godotenv"
	"homeplants-api/internal/config"
	"homeplants-api/internal/database"
	"homeplants-api/internal/handlers"
)

func main() {

    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, relying on system environment variables")
    }

    
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal("Failed to loag config: %v", err)
    }

    db, err := database.NewDatabase(cfg)
    if err != nil {
        log.Fatal("Failed to connect to Database: %v", err)
    }
    defer db.Close()

    roomHandler := handlers.NewRoomHandler(db)
    plantHandler := handlers.NewPlantHandler(db)

   
    http.HandleFunc("GET /rooms", roomHandler.GetRooms)
    http.HandleFunc("GET /rooms/{id}", roomHandler.GetRoomByID)
    http.HandleFunc("GET /rooms/{id}/plants", roomHandler.GetPlantsByRoom)

    
    // http.HandleFunc("GET /plants", handlers.GetPlants)
    // http.HandleFunc("GET /plants/{id}", handlers.GetPlantByID)
    // http.HandleFunc("POST /plants", handlers.CreatePlant)
    http.HandleFunc("POST /plants", plantHandler.CreatePlant)

    
    fmt.Println("Server starting on :8080 ...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}