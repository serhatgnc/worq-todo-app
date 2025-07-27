import { fireEvent, render, screen } from "@testing-library/react";
import { TodoApp } from "./TodoApp";

describe("TodoApp Acceptance Tests", () => {
  test("Given empty list, when I add 'Buy some milk', then it appears in the list", () => {
    render(<TodoApp />);

    const inputElement = screen.getByPlaceholderText(
      "Add a new todo"
    ) as HTMLInputElement;
    const addButton = screen.getByRole("button", { name: "Add" });

    fireEvent.change(inputElement, { target: { value: "Buy some milk" } });

    fireEvent.click(addButton);

    expect(screen.getByText("Buy some milk")).toBeInTheDocument();
  });
});
