package controller

import (
	"database/sql"
	"encoding/json"
	"example/golangserverproject/server/model"
	"example/golangserverproject/server/utils"
	"fmt"
	"net/http"
)

func CreateSqlData(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	query := "INSERT INTO tests(StudentName, Grade) VALUES (?, ?)"
	var student model.Students

	err := utils.HandleRequestBody(writer, request, &student)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return err
	}

	statement, err := database.Prepare(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return err
	}

	dataResult, err := statement.Exec(student.StudentName, student.Grade)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	response := map[string]interface{}{
		"message":      "Data inserted successfully",
		"inserted_row": dataResult,
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		return fmt.Errorf("error encoding while inserting student data: %v", err)
	}

	return nil
}
