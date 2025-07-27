package models

// todo item
type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// request body for creating a todo
type TodoRequest struct {
	Text string `json:"text"`
}