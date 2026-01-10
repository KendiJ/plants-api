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

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	
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

	
	if err := DB.Ping(); err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	fmt.Println("Success connection to MySQL!")
}

func RunMigrations() {
	
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error loading schima.sql:", err)
	}

	
	_, err = DB.Exec(string(schema))
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
	fmt.Println("Tables created successfully!")
}