# Go Learning Repository

A collection of Go tutorials and a project for learning Go programming.

## Structure

- **tutorial_1/** - Basic Go concepts (main, functions, etc.)
- **tutorial_2/** - Goroutines, channels and generics
- **cmd/api/** - Simple REST API project using Chi router
- **internal/** - Internal packages and utilities

## Getting Started

Prerequisites:
- Go 1.26.2 or higher

Running tutorials:
```bash
go run tutorial_1/main.go
go run tutorial_2/channels2.go
```

Running the API:
```bash
go run cmd/api/main.go
```

## Dependencies

- github.com/go-chi/chi - HTTP router
- github.com/sirupsen/logrus - Logging library
- github.com/gorilla/schema - Form/query parsing

## Usage

Install dependencies:
```bash
go mod tidy
```

Build and run:
```bash
go build -o app ./cmd/api
./app
```
