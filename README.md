# Go Initial HTTP Server

This is a simple Go HTTP server project using the Gin framework for routing, Zap for logging, and Docker for containerization.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Build and Run Locally](#build-and-run-locally)
  - [Run Tests](#run-tests)
  - [Build and Run with Docker](#build-and-run-with-docker)
  - [Stop Docker Container](#stop-docker-container)
- [Project Structure](#project-structure)

## Getting Started

### Prerequisites

Before you begin, ensure you have Go and Docker installed on your machine.

### Build and Run Locally

To build and run the server locally, use the following commands:

```bash
make run
```

The server will start and listen on the specified port. You can access it by navigating to [http://localhost:8080/hello](http://localhost:8080/hello) in your browser or using a tool like `curl`.

### Run Tests

To run tests, use the following command:

```bash
make test
```

### Build and Run with Docker

To build and run the server using Docker, use the following commands:

```bash
make docker-build
make docker-run
```

### Stop Docker Container

To stop the Docker container, use the following command:

```bash
make docker-stop
```

## Project Structure

The project structure is organized as follows:

```plaintext
go-initial-http-server/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   └── hello_handler.go
│   └── logger/
│       └── logger.go
├── Makefile
├── Dockerfile
├── go.mod
├── go.sum
└── config.json
```

- `cmd/my_server/main.go`: Main entry point for the application.
- `internal/config/config.go`: Configuration handling using JSON.
- `internal/handler/hello_handler.go`: Example HTTP handler.
- `internal/logger/logger.go`: Logging setup using Zap.
- `Makefile`: Makefile for build, run, test, and Docker commands.
- `Dockerfile`: Dockerfile for containerization.
- `config.json`: Configuration file for the application.
