import axios from "axios";

export interface Todo {
  id: string;
  text: string;
}

export interface CreateTodoRequest {
  text: string;
}

const api = axios.create({
  baseURL:
    process.env.NODE_ENV === "test"
      ? "http://127.0.0.1:1234" // Pact mock server
      : "/api", // Real backend server
  headers: {
    "Content-Type": "application/json",
  },
});

export const todoService = {
  async addTodo(text: string): Promise<Todo> {
    const response = await api.post("/todos", { text });
    return response.data;
  },

  async getTodos(): Promise<Todo[]> {
    const response = await api.get("/todos");
    return response.data;
  },
};
