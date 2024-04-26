package connection

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func DatabaseConnection() (*sql.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("SQL_Link"))
	if err != nil {
		fmt.Println("there was an error connecting to database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return nil, err
	}

	return db, nil
}
