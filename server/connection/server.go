package connection

import (
	"database/sql"
	"example/golangserverproject/server/controller"
	"net/http"

	"github.com/gorilla/mux"
)

// server type struct of mux gorilla router
type Server struct {
	*mux.Router
}

// function MainServer passes database parameter with a server pointer return
func MainServer(database *sql.DB) *Server {
	server := &Server{
		Router: mux.NewRouter(),
	}

	server.routes(database)
	return server
}

// handles all routing
func (server *Server) routes(database *sql.DB) {
	server.HandleFunc("/getStudents", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.GetSQLData(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("GET")

	server.HandleFunc("/getStudents/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.GetSQLDataByID(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("GET")

	server.HandleFunc("/createStudent", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.CreateSqlData(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("POST")

	server.HandleFunc("/updateStudentByName/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.UpdateSqlDataByName(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("PUT")

	server.HandleFunc("/updateStudentByGrade/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.UpdateSqlDataByGrade(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("PUT")

	server.HandleFunc("/updateStudent/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.UpdateSqlData(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("PUT")

	server.HandleFunc("/deleteStudent/{id}", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.DeleteStudentById(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("DELETE")

	server.HandleFunc("/deleteStudent", func(writer http.ResponseWriter, request *http.Request) {
		err := controller.DeleteStudent(database, writer, request)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
	}).Methods("DELETE")
}
