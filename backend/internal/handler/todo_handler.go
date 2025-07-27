package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"worq-todo-api/internal/models"
	"worq-todo-api/internal/repository"

	"github.com/google/uuid"
)

func TodoHandler(repo repository.TodoRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handleCreateTodo(w, r, repo)
		case http.MethodGet:
			handleGetTodos(w, r, repo)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleCreateTodo(w http.ResponseWriter, r *http.Request, repo repository.TodoRepository) {
	// Parse JSON request
	var todoReq models.TodoRequest
	if err := json.NewDecoder(r.Body).Decode(&todoReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new todo
	todo := models.Todo{
		ID:   uuid.New().String(),
		Text: strings.TrimSpace(todoReq.Text),
	}

	// Save to database via repository
	if err := repo.Create(context.Background(), &todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

func handleGetTodos(w http.ResponseWriter, _ *http.Request, repo repository.TodoRepository) {
	// Get all todos from repository
	todos, err := repo.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}