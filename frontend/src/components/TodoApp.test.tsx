import { fireEvent, render, screen } from "@testing-library/react";
import { TodoApp } from "./TodoApp";

describe("TodoApp Integration Tests", () => {
  test("should render TodoApp component", () => {
    render(<TodoApp />);
    const todoAppElement = screen.getByTestId("todo-app");
    expect(todoAppElement).toBeInTheDocument();
  });

  test("should render input and add button", () => {
    render(<TodoApp />);

    const inputElement = screen.getByPlaceholderText("Add a new todo");
    const addButton = screen.getByRole("button", { name: "Add" });

    expect(inputElement).toBeInTheDocument();
    expect(addButton).toBeInTheDocument();
  });

  test("should start with empty input value", () => {
    render(<TodoApp />);

    const inputElement = screen.getByPlaceholderText(
      "Add a new todo"
    ) as HTMLInputElement;

    expect(inputElement.value).toBe("");
  });

  test("should allow user to type in the todo-input", () => {
    render(<TodoApp />);

    const inputElement = screen.getByPlaceholderText(
      "Add a new todo"
    ) as HTMLInputElement;

    fireEvent.change(inputElement, { target: { value: "buy milk" } });

    expect(inputElement.value).toBe("buy milk");
  });
});
