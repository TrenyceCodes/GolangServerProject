package connection

import (
	"database/sql"
	"fmt"
)

func DatabaseConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/testsDB")
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
