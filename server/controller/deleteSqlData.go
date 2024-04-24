package controller

import (
	"database/sql"
	"encoding/json"
	"example/golangserverproject/server/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteStudentById(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	params := mux.Vars(request)

	idString, err := utils.GetParamID(writer, params)
	if err != nil {
		http.Error(writer, "Missing student ID", http.StatusBadRequest)
		return nil
	}

	id, err := utils.ConvertIDStringToIDInt(idString, writer)
	if err != nil {
		http.Error(writer, "Invalid student ID", http.StatusBadRequest)
		return nil
	}

	query := "DELETE FROM tests WHERE ID = ?"

	statement, err := database.Exec(query, id)
	if err != nil {
		http.Error(writer, "There was an error querying the database", http.StatusInternalServerError)
		return nil
	}

	response := map[string]interface{}{
		"message":     "Data deleted successfully",
		"deleted_Row": statement,
	}

	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		return fmt.Errorf("error encoding while deleting student data")
	}

	return nil
}

func DeleteStudent(database *sql.DB, writer http.ResponseWriter, request *http.Request) error {
	query := "DELETE FROM tests"

	statement, err := database.Query(query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return nil
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Println(statement)
	return nil
}
