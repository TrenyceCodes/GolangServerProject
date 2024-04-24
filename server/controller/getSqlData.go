package controller

import (
	"database/sql"
	"encoding/json"
	"example/golangserverproject/server/model"
	"example/golangserverproject/server/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSQLData(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	rows, err := database.Query("SELECT * FROM tests")
	if err != nil {
		return fmt.Errorf("error querying data from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var student model.Students
		if scanErr := rows.Scan(&student.ID, &student.StudentName, &student.Grade); scanErr != nil {
			return fmt.Errorf("error scanning row: %v", scanErr)
		}

		writer.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(writer).Encode(student); err != nil {
			return fmt.Errorf("error encoding student: %v", err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating through rows: %v", err)
	}

	return nil
}

func GetSQLDataByID(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	params := mux.Vars(request)

	idString, err := utils.GetParamID(writer, params)
	if err != nil {
		http.Error(writer, "Missing student ID", http.StatusBadRequest)
		return nil
	}

	ID, err := utils.ConvertIDStringToIDInt(idString, writer)
	if err != nil {
		http.Error(writer, "Invalid student ID", http.StatusBadRequest)
		return nil
	}

	query := "SELECT * FROM tests WHERE ID = ?"
	rows := database.QueryRow(query, ID)

	var student model.Students
	if err := rows.Scan(&student.ID, &student.StudentName, &student.Grade); err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(writer, request)
			return nil
		}
		http.Error(writer, "There was an error scanning for a specific student", http.StatusInternalServerError)
		return fmt.Errorf("error scanning row for specific student: %v", err)
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(student); err != nil {
		return fmt.Errorf("error encoding student: %v", err)
	}

	return nil
}
