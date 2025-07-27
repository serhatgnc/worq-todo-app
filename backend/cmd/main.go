package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"worq-todo-api/internal/database"
	"worq-todo-api/internal/repository"
	"worq-todo-api/internal/server"
)

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		// Continue to next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// MongoDB connection
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	
	// Connect to MongoDB
	client, err := database.NewMongoClient(mongoURI)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	
	// Test connection
	if err := database.PingDatabase(client); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}
	fmt.Println("✅ Connected to MongoDB successfully")

	// Clean up on exit
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}()

	// Create repository
	db := client.Database("worq_todo")
	repo := repository.NewMongoTodoRepository(db)

	// Create HTTP server
	httpServer := server.NewServer(repo)
	serverWithCORS := corsMiddleware(httpServer)

	// Server configuration
	port := getEnv("PORT", "8080")
	
	fmt.Printf("🚀 Server starting on port %s\n", port)
	fmt.Printf("🌐 CORS enabled for: http://localhost:3000\n")
	fmt.Printf("📋 Health check: http://localhost:%s/health\n", port)
	fmt.Printf("📝 Todos API: http://localhost:%s/todos\n", port)

	// Start server
	if err := http.ListenAndServe(":"+port, serverWithCORS); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// getEnv gets environment variable with fallback
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 