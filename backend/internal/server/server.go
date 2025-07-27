package server

import (
	"net/http"
	"worq-todo-api/internal/handler"
	"worq-todo-api/internal/repository"
)

func NewServer(repo repository.TodoRepository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthHandler)

	mux.HandleFunc("/todos", handler.TodoHandler(repo))

	return mux
}