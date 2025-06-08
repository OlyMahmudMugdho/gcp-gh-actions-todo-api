package main

import (
	"gcp-gh-actions-todo-api/config"
	"gcp-gh-actions-todo-api/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	router := routes.RegisterRoutes()
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", router)
}
