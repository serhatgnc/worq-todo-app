package repository

import (
	"context"
	"testing"
	"time"

	"worq-todo-api/internal/models"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) *mongo.Database {
	// Connect to test database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	// Use test database
	db := client.Database("worq_todo_test")
	
	// Clean up after test
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.Drop(ctx)
		client.Disconnect(ctx)
	})

	return db
}

func TestMongoTodoRepository(t *testing.T) {
	t.Run("should create todo successfully", func(t *testing.T) {
		// ARRANGE
		db := setupTestDB(t)
		repo := NewMongoTodoRepository(db)
		
		todo := &models.Todo{
			ID:   "550e8400-e29b-41d4-a716-446655440000",
			Text: "Buy some milk",
		}

		// ACT
		err := repo.Create(context.Background(), todo)

		// ASSERT
		assert.NoError(t, err)
	})

	t.Run("should get all todos successfully", func(t *testing.T) {
		// ARRANGE
		db := setupTestDB(t)
		repo := NewMongoTodoRepository(db)
		
		// Create test data
		todo1 := &models.Todo{
			ID:   "550e8400-1111-41d4-a716-446655440000",
			Text: "Todo 1",
		}
		todo2 := &models.Todo{
			ID:   "550e8400-2222-41d4-a716-446655440000", 
			Text: "Todo 2",
		}
		
		repo.Create(context.Background(), todo1)
		repo.Create(context.Background(), todo2)

		// ACT
		todos, err := repo.GetAll(context.Background())

		// ASSERT
		assert.NoError(t, err)
		assert.Len(t, todos, 2)
		assert.Equal(t, "Todo 1", todos[0].Text)
		assert.Equal(t, "Todo 2", todos[1].Text)
	})
} 