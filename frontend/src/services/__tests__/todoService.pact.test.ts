import { PactV3, MatchersV3 } from "@pact-foundation/pact";
import path from "path";
import { todoService } from "../todoService";

describe("TodoService Pact Tests", () => {
  const provider = new PactV3({
    consumer: "TodoFrontend",
    provider: "TodoAPI",
    port: 1234,
    logLevel: "info",
    dir: path.resolve(process.cwd(), "pacts"),
  });

  describe("addTodo", () => {
    test("should create todo successfully", async () => {
      provider
        .given("server is healthy")
        .uponReceiving("a request to create todo")
        .withRequest({
          method: "POST",
          path: "/todos",
          headers: {
            "Content-Type": "application/json",
          },
          body: {
            text: "Buy some milk",
          },
        })
        .willRespondWith({
          status: 201,
          headers: {
            "Content-Type": "application/json",
          },
          body: {
            id: MatchersV3.uuid("550e8400-e29b-41d4-a716-446655440000"),
            text: "Buy some milk",
          },
        });

      return provider.executeTest(async () => {
        const todo = await todoService.addTodo("Buy some milk");

        expect(todo.id).toBeDefined();
        expect(todo.text).toBe("Buy some milk");
      });
    });
  });
});
