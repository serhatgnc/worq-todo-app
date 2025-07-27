package database

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	t.Run("should connect to MongoDB successfully", func(t *testing.T) {
		// ARRANGE: MongoDB connection string (use local for testing)
		mongoURI := "mongodb://localhost:27017"
		
		// ACT: create db connection
		client, err := NewMongoClient(mongoURI)
		
		// ASSERT: connection should be successful
		assert.NoError(t, err)
		assert.NotNil(t, client)
		
		// Cleanup
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	})
	
	t.Run("should ping database successfully", func(t *testing.T) {
		// ARRANGE
		mongoURI := "mongodb://localhost:27017"
		client, _ := NewMongoClient(mongoURI)
		
		// ACT: Database ping
		err := PingDatabase(client)
		
		// ASSERT: ping should be successful
		assert.NoError(t, err)
		
		// Cleanup
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	})
} 