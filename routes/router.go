package routes

import (
	"gcp-gh-actions-todo-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	return r
}
