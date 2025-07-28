import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { TodoApp } from "./TodoApp";
import { todoService } from "../services/todoService";

jest.mock("../services/todoService", () => ({
  todoService: {
    addTodo: jest.fn(),
    getTodos: jest.fn(),
  },
}));

const mockAddTodo = todoService.addTodo as jest.MockedFunction<
  typeof todoService.addTodo
>;
const mockGetTodos = todoService.getTodos as jest.MockedFunction<
  typeof todoService.getTodos
>;

describe("TodoApp Acceptance Tests", () => {
  beforeEach(() => {
    jest.clearAllMocks();
    mockAddTodo.mockResolvedValue({
      id: "1",
      text: "Buy some milk",
    });

    mockGetTodos.mockResolvedValue([]);
  });
  test("Given empty list, when I add 'Buy some milk', then it appears in the list", async () => {
    render(<TodoApp />);

    const inputElement = screen.getByPlaceholderText(
      "Add a new todo"
    ) as HTMLInputElement;
    const addButton = screen.getByRole("button", { name: "Add" });

    fireEvent.change(inputElement, { target: { value: "Buy some milk" } });
    fireEvent.click(addButton);

    // Wait for the todo to appear
    await waitFor(() => {
      expect(screen.getByText("Buy some milk")).toBeInTheDocument();
    });

    // Verify the service was called
    expect(mockAddTodo).toHaveBeenCalledWith("Buy some milk");
  });
});
