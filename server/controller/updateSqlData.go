package controller

import (
	"database/sql"
	"encoding/json"
	"example/golangserverproject/server/model"
	"example/golangserverproject/server/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateSqlData(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	params := mux.Vars(request)
	var student model.Students
	query := "UPDATE tests SET StudentName = ?, Grade = ? WHERE ID = ?"

	err := utils.HandleRequestBody(writer, request, &student)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	idString, err := utils.GetParamID(writer, params)
	if err != nil {
		http.Error(writer, "Missing student id", http.StatusBadRequest)
	}

	ID, err := utils.ConvertIDStringToIDInt(idString, writer)
	if err != nil {
		http.Error(writer, "Invalid id", http.StatusBadRequest)
		return nil
	}

	statement, err := database.Exec(query, student.StudentName, student.Grade, ID)
	if err != nil {
		http.Error(writer, "There was an issue with the statement", http.StatusInternalServerError)
		return nil
	}

	response := map[string]interface{}{
		"message":   "Data successfully updated",
		"status":    http.StatusOK,
		"id":        ID,
		"statement": statement,
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return nil
}

func UpdateSqlDataByName(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	params := mux.Vars(request)
	var student model.Students
	query := "UPDATE tests SET StudentName = ? WHERE ID = ? "

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	if err := json.Unmarshal(body, &student); err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
		return err
	}

	idString, ok := params["id"]
	if !ok {
		http.Error(writer, "missing student id", http.StatusBadRequest)
		return nil
	}

	ID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	statement, err := database.Exec(query, student.StudentName, ID)
	if err != nil {
		http.Error(writer, "There was an issue with the statement", http.StatusInternalServerError)
		return nil
	}

	response := map[string]interface{}{
		"message":   "Data successfully updated",
		"status":    http.StatusOK,
		"id":        ID,
		"statement": statement,
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return nil
}

func UpdateSqlDataByGrade(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	params := mux.Vars(request)
	var student model.Students
	query := "UPDATE tests SET Grade = ? WHERE ID = ? "

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	if err := json.Unmarshal(body, &student); err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
		return err
	}

	idString, ok := params["id"]
	if !ok {
		http.Error(writer, "missing student id", http.StatusBadRequest)
		return nil
	}

	ID, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	statement, err := database.Exec(query, student.Grade, ID)
	if err != nil {
		http.Error(writer, "There was an issue with the statement", http.StatusInternalServerError)
		return nil
	}

	response := map[string]interface{}{
		"message":   "Data successfully updated",
		"status":    http.StatusOK,
		"id":        ID,
		"statement": statement,
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return nil
}
