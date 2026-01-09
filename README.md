SnippetBox

SnippetBox is a full-stack web application written in Go that allows users to create, view, and manage text snippets — similar in concept to Pastebin.
This project is based on the Lets Go book and serves as a practical demonstration of building a production-ready web application using idiomatic Go.

The application follows Go best practices for project structure, dependency management, security, and web development.

Features

Create and view text snippets

Automatically expire snippets after a set period

Server-side rendered HTML using Go templates

Secure user input handling

Structured logging and error handling

Clean separation of concerns across application layers

Project Structure

The project is organized according to commonly accepted Go conventions:

.
├── cmd/
│   └── web/
│       └── main.go
├── internal/
│   ├── models/
│   ├── validator/
│   └── ...
├── ui/
│   ├── html/
│   └── static/
└── go.mod

cmd/

Contains the application-specific entry point for the executable programs.

cmd/web holds the web server configuration, routing, middleware, and startup logic.

This separation allows additional executables (e.g., background workers or APIs) to be added cleanly in the future.

internal/

Houses non-application-specific business logic that should not be imported by external projects.

models — database models and data access logic

validator — form validation helpers

Shared utilities and helpers used across the application

Using the internal directory enforces encapsulation at the compiler level.

ui/

Contains all user interface assets for the web application.

ui/html — Go HTML templates

ui/static — static assets such as CSS and images

This separation keeps presentation concerns isolated from application logic.

Technologies Used

Go — backend language

net/http — HTTP server and routing

html/template — server-side HTML rendering

MySQL (or SQLite, if applicable) — data persistence

Go Modules — dependency management

Screenshots

Below is a screenshot of the SnippetBox homepage:

![SnippetBox Homepage](./ui/static/img/screenshot.png)


(Replace the path above with the actual path to your image.)

Including screenshots helps reviewers quickly understand the UI and overall functionality without running the app locally.

What This Project Demonstrates

Writing idiomatic, maintainable Go

Structuring a real-world Go web application

Secure handling of user input and errors

Separation of concerns between routing, business logic, and presentation

Readiness for extension and production hardening

Future Improvements

User authentication and authorization

Snippet search functionality

Pagination and filtering

Deployment configuration (Docker / cloud hosting)
