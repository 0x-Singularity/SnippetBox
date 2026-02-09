# SnippetBox

SnippetBox is a **full-stack web application written in Go** that allows users to create, view, and manage text snippets — similar in concept to Pastebin.

This project is based on the *Lets Go* book and serves as a practical demonstration of building a production-ready web application using idiomatic Go. It emphasizes clean architecture, maintainability, and best practices commonly used in real-world Go applications.

---

## Features

- Create and view text snippets
- Automatically expire snippets after a configurable period
- Server-side rendered HTML using Go templates
- Secure handling of user input and errors
- Structured logging
- Clear separation of concerns across application layers

---

## Project Structure

The project follows standard Go project layout conventions:

### `cmd/`

Contains the **application-specific entry point(s)**.

- `cmd/web` holds the web server startup logic, routing, middleware, and configuration.
- This structure makes it easy to add additional executables in the future (e.g., APIs, background workers).

### `internal/`

Contains **core business logic and shared application code** that should not be imported by external projects.

- `models` — database models and data access logic
- `validator` — form validation helpers
- Additional utilities and helpers used across the application

Using the `internal` directory enforces encapsulation at the compiler level.

### `ui/`

Contains all **user interface assets** for the web application.

- `ui/html` — Go HTML templates used for server-side rendering
- `ui/static` — static assets such as CSS and images

This keeps presentation concerns separate from application logic.

---

## Technologies Used

- **Go**
- **net/http**
- **html/template**
- **Go Modules**
- **MySQL** (or SQLite, if applicable)

---

## SnippetBox View
![SnippetBoxView](ui/static/img/SnippetBoxView.)
Testing new PAT
