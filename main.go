package main

import (
	"example/golangserverproject/server/connection"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database, err := connection.DatabaseConnection()

	if err != nil {
		fmt.Println("Error connecting to database", err)
		return
	}
	defer database.Close()

	if database != nil {
		fmt.Println("Database is connected")
	}

	// Start HTTP server
	server := connection.MainServer(database)

	fmt.Println("Server started on localhost:3001")
	if err := http.ListenAndServe("localhost:3001", server); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
