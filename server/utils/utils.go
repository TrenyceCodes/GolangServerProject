package utils

import (
	"encoding/json"
	"example/golangserverproject/server/model"
	"io"
	"net/http"
	"strconv"
)

func ConvertIDStringToIDInt(id string, writer http.ResponseWriter) (int, error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(writer, "Invalid id", http.StatusBadRequest)
		return 0, nil
	}

	return ID, nil
}

func HandleRequestBody(writer http.ResponseWriter, request *http.Request, student *model.Students) error {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return nil
	}

	if err := json.Unmarshal(body, &student); err != nil {
		http.Error(writer, err.Error(), http.StatusBadGateway)
		return nil
	}

	return nil
}

func GetParamID(writer http.ResponseWriter, params map[string]string) (string, error) {

	idString, ok := params["id"]
	if !ok {
		http.Error(writer, "Missing student id", http.StatusBadRequest)
	}

	return idString, nil
}
