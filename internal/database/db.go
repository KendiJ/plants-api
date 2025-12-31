package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB () {

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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", user, pass, host, port, name)


	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err) // this crashes app instantly
	}

	// test the connection
	if err := DB.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("Success connection to MySQL!")
}

func RunMigrations() {
	// Read the file from the project root
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error loading schima.sql:", err)
	}

	//run sql
	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
	fmt.Println("Tables created successfully!")
}