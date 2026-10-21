# Product Management API (Go)

A lightweight, high-performance RESTful API for managing products built with Go using only its standard library (`net/http`). This project demonstrates standard Go project layout, concurrency safety using `sync.RWMutex`, and clean multi-layer architecture.

## 🚀 Features

- **Full CRUD Support**:
  - `POST /products` - Create a new product
  - `GET /products` - List all products
  - `GET /products/{id}` - Get a specific product by ID
  - `PUT /products/{id}` - Update an existing product
  - `DELETE /products/{id}` - Delete a product
- **Clean Architecture**: Separated into `models`, `repository`, and `handlers`.
- **Thread-Safe**: Uses `sync.RWMutex` to handle concurrent HTTP requests safely in memory.
- **Zero External Dependencies**: Built entirely with Go native modules.

## 🛠️ Project Structure

```text
minha-api-produtos/
├── go.mod
├── models/
│   └── product.go
├── repository/
│   └── product.go
├── handlers/
│   └── product.go
├── main.go
└── README.md
```

## ⚙️ Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed.

### Running the Application

1. Clone or download this repository.
2. Open your terminal in the project root directory.
3. Run the application:

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

## 🧪 Example Request

**Create a product:**

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mechanical Keyboard",
    "description": "RGB Backlit Switches",
    "price": 99.99
  }'
```