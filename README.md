# Simple TCP Chat Application in Go

This is a simple TCP-based chat application written in Go. It includes a **server** and a **client**, allowing two-way communication over a TCP connection.

## Features

- Text-based real-time chat via terminal
- Uses Go's `net` package for TCP connections
- Handles concurrent message receiving with goroutines

## Files

- `server.go`: Starts a TCP server that listens on port `8080` and accepts a single client connection.
- `client.go`: Connects to the server and allows the user to send and receive messages.

## How It Works

- The server waits for a client to connect.
- Both server and client run a goroutine to handle incoming messages.
- Messages typed into the terminal are sent to the other party.
- The server exits cleanly when `/quit` is typed.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) (v1.18 or newer recommended)

### Running the Server

```bash
go run server/server.go
```

### Running the Clinet (in another terminal)

```bash
go run client/client.go
