package handlers

import (
	"context"
	"encoding/json"
	"gcp-gh-actions-todo-api/config"
	"gcp-gh-actions-todo-api/models"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, title, completed FROM todos")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t models.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := config.DB.QueryRow(context.Background(),
		"INSERT INTO todos(title, completed) VALUES($1, $2) RETURNING id",
		t.Title, t.Completed,
	).Scan(&t.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
