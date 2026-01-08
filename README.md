🎬 Movie CRUD API (Go + Gorilla Mux)

A simple RESTful CRUD API built using Go (Golang) and Gorilla Mux to learn HTTP servers, routing, handlers, and JSON handling in Go.

This project uses in-memory storage and is intended for learning and practice, not production use.

🚀 Features

Create a movie

Get all movies

Get a movie by ID

Update a movie by ID

Delete a movie by ID

JSON-based request and response handling

RESTful routing using Gorilla Mux

🧱 Tech Stack

Go

net/http

github.com/gorilla/mux

encoding/json

📁 Project Structure
.
├── go.mod
├── go.sum
└── main.go

📦 Data Model
Movie
{
  "id": "1",
  "isbn": "756453",
  "title": "xyz",
  "director": {
    "firstname": "a",
    "lastname": "b"
  }
}

▶️ Getting Started
1. Clone the repository
git clone <repository-url>
cd <repository-name>

2. Install dependencies
go mod tidy

3. Run the server
go run main.go


The server will start at:

http://localhost:8080

🔌 API Endpoints
Get all movies
GET /movies

Get movie by ID
GET /movies/{id}

Create a movie
POST /movies


Request Body

{
  "isbn": "123456",
  "title": "Inception",
  "director": {
    "firstname": "Christopher",
    "lastname": "Nolan"
  }
}

Update a movie
PUT /movies/{id}

Delete a movie
DELETE /movies/{id}

🧠 Request Flow
Client
  ↓
HTTP Server
  ↓
Gorilla Mux Router
  ↓
Handler Functions
  ↓
In-Memory Data Store
  ↓
JSON Response

⚠️ Limitations

No database (data resets on restart)

No validation or detailed error handling

Not concurrency-safe

Random ID generation

These limitations are intentional to keep the project focused on core Go and HTTP concepts.

🛣️ Possible Improvements

Add proper HTTP status codes

Add request validation

Add middleware (logging, recovery)

Introduce a database (PostgreSQL / MongoDB)

Refactor into service and repository layers

Add concurrency safety

👤 Author

Lakshay Singhal
