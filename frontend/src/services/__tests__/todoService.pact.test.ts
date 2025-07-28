import { PactV3, MatchersV3 } from "@pact-foundation/pact";
import path from "path";
import { todoService } from "../todoService";

describe("TodoService Pact Tests", () => {
  const provider = new PactV3({
    consumer: "TodoFrontend",
    provider: "TodoAPI",
    port: 1234,
    host: "127.0.0.1",
    logLevel: "info",
    dir: path.resolve(process.cwd(), "pacts"),
  });

  describe("addTodo", () => {
    test("should create todo successfully", async () => {
      await provider
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
            id: String(MatchersV3.uuid()),
            text: "Buy some milk",
          },
        })
        .executeTest(async () => {
          const todo = await todoService.addTodo("Buy some milk");

          expect(todo.id).toBeDefined();
          expect(todo.text).toBe("Buy some milk");
        });
    });
  });
});
