import { test, expect } from "@playwright/test";

test.describe("Todo App E2E", () => {
  test.beforeEach(async ({ page }) => {
    // 🧹 Clear database before each test
    await page.request.delete("http://localhost:8080/todos");

    // Navigate to the app
    await page.goto("/");
  });

  test("should display todo app", async ({ page }) => {
    // Check title
    await expect(page.locator("h1")).toContainText("WORQ Todo App");

    // Check input and button
    await expect(page.getByTestId("todo-input")).toBeVisible();
    await expect(page.getByTestId("todo-add-button")).toBeVisible();
  });

  test("should add and display new todo", async ({ page }) => {
    // Add a new todo
    const timestamp = Date.now();
    await page.getByTestId("todo-input").fill(`buy some milk ${timestamp}`);
    await page.getByTestId("todo-add-button").click();

    // Check if todo appears in the list
    await expect(page.getByText(`buy some milk ${timestamp}`)).toBeVisible();

    // Check if input is cleared
    await expect(page.getByTestId("todo-input")).toHaveValue("");
  });

  test("should persist todos after page refresh", async ({ page }) => {
    // Add a todo
    const timestamp = Date.now();
    await page.getByTestId("todo-input").fill(`Persistent todo ${timestamp}`);
    await page.getByTestId("todo-add-button").click();

    // Refresh the page
    await page.reload();

    // Check if todo is still there
    await expect(page.getByText(`Persistent todo ${timestamp}`)).toBeVisible();
  });

  test("should handle multiple todos", async ({ page }) => {
    const timestamp = Date.now();
    const todos = [
      `First todo ${timestamp}`,
      `Second todo ${timestamp}`,
      `Third todo ${timestamp}`,
    ];

    for (const todo of todos) {
      await page.getByTestId("todo-input").fill(todo);
      await page.getByTestId("todo-add-button").click();

      // 🔧 Wait for todo to appear before adding next one
      await expect(page.getByText(todo)).toBeVisible();
    }

    for (const todo of todos) {
      await expect(page.getByText(todo)).toBeVisible();
    }
  });

  test("should not add empty todos", async ({ page }) => {
    // Try to add empty todo
    await page.getByTestId("todo-add-button").click();

    // Check if no empty todo is added
    await expect(page.getByTestId("todo-list").locator("li")).toHaveCount(0);
  });
});
