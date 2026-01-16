# Go Web Service with Gin: Key Learnings

This document outlines key concepts and best practices observed while building a simple web service using Go and the Gin framework.

## Core Go Concepts

*   **Package `main`**: In Go, an executable program (as opposed to a reusable library) must be in the `main` package. The `main` function within this package serves as the entry point for the application.

*   **Struct Tags**: Struct tags are small pieces of metadata attached to struct fields, enclosed in backticks (``). They are used to provide instructions to other packages, such as the `encoding/json` package.
    *   **Example**: `json:"artist"` tells the JSON encoder to use "artist" as the key in the JSON output, rather than the capitalized struct field name "Artist". This is idiomatic for JSON, which typically uses camelCase or snake_case for keys.

*   **Passing Functions as Arguments**: In Go, functions are first-class citizens. This means they can be passed as arguments to other functions. When setting up HTTP handlers in Gin, you pass the function itself (e.g., `getBooks`), not the result of calling the function (e.g., `getBooks()`)

## Gin Framework Specifics

*   **`gin.Context`**: This is a critical part of Gin. It's a struct that holds all the information about an incoming HTTP request, including request parameters, headers, and the request body. It also provides methods for writing the HTTP response, such as serializing data to JSON.
    > **Note**: This is specific to the Gin framework and is distinct from Go's built-in `context` package, which is used for managing deadlines, cancellation signals, and other request-scoped values across APIs and processes.

*   **Routing and URL Parameters**: Gin provides a simple and expressive way to define API routes. You can define routes with parameters, which can be retrieved from the `gin.Context`.
    *   **Example**: In the route `/books/:id`, the `:id` part is a placeholder for a dynamic value. You can access this value using `c.Param("id")`.

*   **Data Binding**: Gin's `BindJSON` method is a convenient way to parse JSON from the request body and bind it to a Go struct. It automatically handles the unmarshaling process and can return an error if the request body is not valid JSON or does not match the struct's shape.

*   **JSON Responses**: The `IndentedJSON` method on `gin.Context` serializes a Go struct or map into a nicely formatted JSON string and sends it as the HTTP response. It also allows you to set the HTTP status code. Using `c.IndentedJSON` is helpful for development as it makes the JSON output more readable. For production, `c.JSON` is more efficient as it doesn't add whitespace.

*   **Returning HTTP Status Codes**: It is important to return appropriate HTTP status codes to the client. Gin provides constants for standard HTTP statuses, such as `http.StatusOK` (200), `http.StatusCreated` (201), and `http.StatusNotFound` (404).

## Sample Data

```
books=# SELECT * FROM book;
 id |                         title                          |      author      | price
----+--------------------------------------------------------+------------------+-------
  1 | Clean Code                                             | Uncle Bob        | 21.12
  2 | The Pragmatic Programmer                               | David Thomas     | 31.99
  3 | Designing Data-Intensive Applications                  | Martin Kleppmann | 37.00
  4 | Code Complete                                          | Steve McConnell  | 40.01
  5 | Functional Design: Principles, Patterns, and Practices | Uncle Bob        | 49.99
(5 rows)
```

## Example `curl` Commands

**Get all albums:**
```bash
curl http://localhost:8080/albums
```

**Add a new album:**
```bash
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "5","title": "Functional Design: Principles, Patterns, and Practices","artist": "Uncle Bob","price": 49.99}'
```

**Get a specific album:**
```bash
curl http://localhost:8080/albums/2
```
