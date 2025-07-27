import { useEffect, useState } from "react";
import { Todo, todoService } from "../services/todoService";

export const TodoApp = () => {
  const [inputValue, setInputValue] = useState("");
  const [todos, setTodos] = useState<Todo[]>([]);
  const [error, setError] = useState<string | null>(null);

  const handleAddTodo = async () => {
    const trimmedValue = inputValue.trim();

    if (!trimmedValue) {
      return;
    }

    try {
      setError(null);
      const newTodo = await todoService.addTodo(trimmedValue);
      setTodos([...todos, newTodo]);
      setInputValue("");
    } catch (err) {
      setError("Failed to add todo");
      console.error("Error adding todo:", err);
    }
  };

  const loadTodos = async () => {
    try {
      setError(null);
      const fetchedTodos = await todoService.getTodos();
      setTodos(fetchedTodos);
    } catch (err) {
      setError("Failed to load todos");
      console.error("Error loading todos:", err);
    }
  };

  useEffect(() => {
    loadTodos();
  }, []);

  return (
    <div data-testid="todo-app" className="max-w-md mx-auto p-6">
      <h1 className="text-2xl font-bold text-gray-800 mb-6">WORQ Todo App</h1>

      {error && (
        <div className="mb-4 p-3 bg-red-100 border border-red-400 text-red-700 rounded">
          {error}
        </div>
      )}

      <div className="flex gap-2 mb-4">
        <input
          data-testid="todo-input"
          placeholder="Add a new todo"
          className="flex-1 px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
        />
        <button
          data-testid="todo-add-button"
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
          onClick={handleAddTodo}
        >
          Add
        </button>
      </div>

      <ul className="space-y-2">
        {todos.map((todo) => (
          <li key={todo.id} className="p-3 bg-gray-100 rounded shadow-sm">
            {todo.text}
          </li>
        ))}
      </ul>

      {todos.length === 0 && (
        <div className="text-center py-8 text-gray-500">
          No todos yet. Add your first todo above!
        </div>
      )}
    </div>
  );
};
