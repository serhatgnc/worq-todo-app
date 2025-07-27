package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"worq-todo-api/internal/models"
	"worq-todo-api/internal/repository"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestServer(t *testing.T) (*http.ServeMux, repository.TodoRepository) {
	// Setup test database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	db := client.Database("worq_todo_integration_test")
	
	// Clean up after test
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.Drop(ctx)
		client.Disconnect(ctx)
	})

	// Create repository
	repo := repository.NewMongoTodoRepository(db)
	
	// Create server with repository
	server := NewServer(repo)
	
	return server, repo
}

func TestServerIntegration(t *testing.T) {
	t.Run("should create todo and persist to database", func(t *testing.T) {
		// ARRANGE
		server, repo := setupTestServer(t)
		
		todoReq := models.TodoRequest{Text: "Buy some milk"}
		jsonData, _ := json.Marshal(todoReq)
		
		req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		recorder := httptest.NewRecorder()

		// ACT: Call API
		server.ServeHTTP(recorder, req)

		// ASSERT: API Response
		assert.Equal(t, http.StatusCreated, recorder.Code)
		
		var response models.Todo
		json.Unmarshal(recorder.Body.Bytes(), &response)
		assert.NotEmpty(t, response.ID)
		assert.Equal(t, "Buy some milk", response.Text)

		// ASSERT: Database Persistence
		todos, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Len(t, todos, 1)
		assert.Equal(t, "Buy some milk", todos[0].Text)
	})

	t.Run("should get all todos from database", func(t *testing.T) {
		// ARRANGE: Create server and add some todos
		server, repo := setupTestServer(t)
		
		// Create todos directly via repository
		todo1 := &models.Todo{ID: "id1", Text: "First todo"}
		todo2 := &models.Todo{ID: "id2", Text: "Second todo"}
		repo.Create(context.Background(), todo1)
		repo.Create(context.Background(), todo2)

		// ACT: Call GET /todos API
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, req)

		// ASSERT: API Response
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		
		var response []*models.Todo
		json.Unmarshal(recorder.Body.Bytes(), &response)
		
		assert.Len(t, response, 2)
		assert.Equal(t, "First todo", response[0].Text)  
		assert.Equal(t, "Second todo", response[1].Text)
	})

	t.Run("should return empty array when no todos exist in database", func(t *testing.T) {
		// ARRANGE: Empty database
		server, _ := setupTestServer(t)

		// ACT: Call GET /todos API
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		recorder := httptest.NewRecorder()
		server.ServeHTTP(recorder, req)

		// ASSERT: Empty array response
		assert.Equal(t, http.StatusOK, recorder.Code)
		
		var response []*models.Todo
		json.Unmarshal(recorder.Body.Bytes(), &response)
		
		assert.Len(t, response, 0)
	})
} 