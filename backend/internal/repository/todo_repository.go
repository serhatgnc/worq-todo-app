package repository

import (
	"context"
	"worq-todo-api/internal/models"
)

// TodoRepository defines the interface for todo data operations
type TodoRepository interface {
	// Create a new todo
	Create(ctx context.Context, todo *models.Todo) error
	
	// Get all todos
	GetAll(ctx context.Context) ([]*models.Todo, error)
} 