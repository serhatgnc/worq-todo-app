package repository

import (
	"context"

	"worq-todo-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoTodoRepository implements TodoRepository interface using MongoDB
type MongoTodoRepository struct {
	collection *mongo.Collection
}

// NewMongoTodoRepository creates a new MongoDB todo repository
func NewMongoTodoRepository(db *mongo.Database) TodoRepository {
	return &MongoTodoRepository{
		collection: db.Collection("todos"),
	}
}

// Create saves a new todo to MongoDB
func (r *MongoTodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	_, err := r.collection.InsertOne(ctx, todo)
	return err
}

// retrieves all todos from MongoDB
func (r *MongoTodoRepository) GetAll(ctx context.Context) ([]*models.Todo, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	todos := make([]*models.Todo, 0)
	
	for cursor.Next(ctx) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return todos, nil
} 