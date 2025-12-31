package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)
var db *sql.DB

func initDB () {

	// locate .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Construct DSN from Environment Variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)


	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // this crashes app instantly
	}

	// test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("Success connection to MySQL!")
}
func main() {
	//initialize DB
	initDB()
	defer db.Close() // close DB when main function closes

	// setup routes
	http.HandleFunc("/plants", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Connecting to DB...")
	})

	// start the server
	fmt.Println("Server starting on : 8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}