package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"worq-todo-api/internal/models"
	"worq-todo-api/internal/repository"

	"github.com/stretchr/testify/assert"
)

type mockTodoRepository struct {
	todos []*models.Todo
}

func (m *mockTodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	m.todos = append(m.todos, todo)
	return nil
}

func (m *mockTodoRepository) GetAll(ctx context.Context) ([]*models.Todo, error) {
	return m.todos, nil
}

func newMockRepository() repository.TodoRepository {
	return &mockTodoRepository{
		todos: make([]*models.Todo, 0),
	}
}

func TestServer(t *testing.T) {
	t.Run("should create HTTP server", func(t *testing.T) {
		// ARRANGE: create server
		mockRepo := newMockRepository()
		server := NewServer(mockRepo)
		
		// ACT: send GET request
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		mockResponse := httptest.NewRecorder()
		
		server.ServeHTTP(mockResponse, req)
		
		// ASSERT: is server returning ok response
		assert.Equal(t, http.StatusOK, mockResponse.Code)
		assert.Equal(t, "OK", mockResponse.Body.String())
	})

	t.Run("should create todo via POST /todos", func(t *testing.T) {
		// ARRANGE
		mockRepo := newMockRepository()
		server := NewServer(mockRepo)
		
		todoData := map[string]string{
			"text": "Buy some milk",
		}
		jsonData, _ := json.Marshal(todoData)

		// ACT: send POST request
		req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		mockResponse := httptest.NewRecorder()
		
		server.ServeHTTP(mockResponse, req)
		
		// ASSERT: success response
		assert.Equal(t, http.StatusCreated, mockResponse.Code)
		
		// response should contain todo data
		var response map[string]interface{}
		json.Unmarshal(mockResponse.Body.Bytes(), &response)
		assert.Equal(t, "Buy some milk", response["text"])
		assert.NotEmpty(t, response["id"])
	})

	t.Run("should generate unique IDs for multiple todos", func(t *testing.T) {
		// ARRANGE
		mockRepo := newMockRepository()
		server := NewServer(mockRepo)
		
		// ACT: create 2 todos
		todoData1 := `{"text": "First todo"}`
		todoData2 := `{"text": "Second todo"}`
		
		// first todo
		req1 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoData1))
		req1.Header.Set("Content-Type", "application/json")
		mockResponse1 := httptest.NewRecorder()
		server.ServeHTTP(mockResponse1, req1)
		
		// second todo
		req2 := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoData2))
		req2.Header.Set("Content-Type", "application/json")  
		mockResponse2 := httptest.NewRecorder()
		server.ServeHTTP(mockResponse2, req2)
		
		// ASSERT: should be 2 different ids
		var todo1, todo2 map[string]interface{}
		json.Unmarshal(mockResponse1.Body.Bytes(), &todo1)
		json.Unmarshal(mockResponse2.Body.Bytes(), &todo2)
		
		assert.NotEqual(t, todo1["id"], todo2["id"])
		
		// should be uuid format
		id1 := todo1["id"].(string)
		id2 := todo2["id"].(string)
		assert.Len(t, id1, 36) // UUID length
		assert.Len(t, id2, 36)
		assert.Contains(t, id1, "-") // UUID has dashes
		assert.Contains(t, id2, "-")
	})

	t.Run("should get all todos via GET /todos", func(t *testing.T) {
		// ARRANGE: Mock repository with some todos
		mockRepo := newMockRepository()
		server := NewServer(mockRepo)
		
		// Create some todos
		todo1 := &models.Todo{ID: "id1", Text: "First todo"}
		todo2 := &models.Todo{ID: "id2", Text: "Second todo"}
		mockRepo.Create(context.Background(), todo1)
		mockRepo.Create(context.Background(), todo2)

		// ACT: send GET request
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		recorder := httptest.NewRecorder()
		
		server.ServeHTTP(recorder, req)
		
		// ASSERT: success response with todos array
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		
		var response []*models.Todo
		json.Unmarshal(recorder.Body.Bytes(), &response)
		
		assert.Len(t, response, 2)
		assert.Equal(t, "First todo", response[0].Text)
		assert.Equal(t, "Second todo", response[1].Text)
		assert.Equal(t, "id1", response[0].ID)
		assert.Equal(t, "id2", response[1].ID)
	})

	t.Run("should return empty array when no todos exist", func(t *testing.T) {
		// ARRANGE: Empty mock repository
		mockRepo := newMockRepository()
		server := NewServer(mockRepo)

		// ACT: send GET request
		req := httptest.NewRequest(http.MethodGet, "/todos", nil)
		recorder := httptest.NewRecorder()
		
		server.ServeHTTP(recorder, req)
		
		// ASSERT: success response with empty array
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
		
		var response []*models.Todo
		json.Unmarshal(recorder.Body.Bytes(), &response)
		
		assert.Len(t, response, 0)
	})
} 