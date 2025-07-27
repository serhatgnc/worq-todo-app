package server

import (
	"context"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
	"worq-todo-api/internal/repository"

	"github.com/pact-foundation/pact-go/v2/provider"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestPactProvider(t *testing.T) {
	contractPath, _ := filepath.Abs("../../../frontend/pacts/TodoFrontend-TodoAPI.json")

	// Setup test database for Pact verification
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	db := client.Database("worq_todo_pact_test")
	
	// Clean up after test
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		db.Drop(ctx)
		client.Disconnect(ctx)
	})
	
	// Create repository and server
	repo := repository.NewMongoTodoRepository(db)
	server := NewServer(repo)
	testServer := httptest.NewServer(server)
	defer testServer.Close()
	
	err = provider.NewVerifier().VerifyProvider(t, provider.VerifyRequest{
		ProviderBaseURL: testServer.URL,
		PactFiles:       []string{contractPath},
	})
	
	if err != nil {
		t.Fatalf("Pact verification failed: %v", err)
	}
}