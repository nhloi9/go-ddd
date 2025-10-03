# DDD Go Example - Tavern Management System

A Domain-Driven Design (DDD) example implementation in Go, demonstrating a tavern management system with customer and product management, order processing, and multiple repository implementations.

## 🏗️ Architecture Overview

This project follows DDD principles with a clean, layered architecture:

```
ddd-go/
├── aggregate/          # Domain aggregates (Customer, Product)
├── domain/            # Domain layer with repositories
│   ├── customer/      # Customer domain
│   └── product/       # Product domain
├── entity/            # Domain entities (Person, Item)
├── services/          # Application services (Tavern, Order)
├── valueobject/       # Value objects (Transaction)
└── race_demo.go       # Concurrency demonstration
```

## 🧱 DDD Layers

### **Aggregates** (`/aggregate`)

- **Customer**: Root aggregate managing customer data, products, and transactions
- **Product**: Product aggregate with inventory management

### **Entities** (`/entity`)

- **Person**: Customer personal information
- **Item**: Product items with pricing

### **Value Objects** (`/valueobject`)

- **Transaction**: Immutable transaction records

### **Domain Services** (`/services`)

- **Tavern**: Main application service orchestrating business operations
- **OrderService**: Handles order processing and customer management

### **Repositories** (`/domain`)

- **CustomerRepository**: Interface for customer data persistence
- **Memory Implementation**: In-memory storage with thread-safe operations
- **MongoDB Implementation**: NoSQL database persistence

## 🚀 Getting Started

### Prerequisites

- Go 1.24.4 or higher
- MongoDB (optional, for mongo repository)

### Installation

```bash
# Clone the repository
git clone https://github.com/nhloi9/ddd-go.git
cd ddd-go

# Download dependencies
go mod download

# Run tests
go test ./...
```

### Usage Example

```go
package main

import (
    "github.com/nhloi9/ddd-go/services"
    "github.com/nhloi9/ddd-go/domain/customer/memory"
)

func main() {
    // Create customer repository
    customerRepo := memory.New(make(map[uuid.UUID]aggregate.Customer))

    // Create order service
    orderService, _ := services.NewOrderService(
        services.WithCustomerRepository(customerRepo),
    )

    // Create tavern with order service
    tavern, _ := services.NewTavern(
        services.WithOrderService(orderService),
    )

    // Use tavern for business operations
    customerID, _ := tavern.AddCustomer("John Doe")
    // ... more operations
}
```

## 🔧 Key Features

### **Thread-Safe Operations**

- Memory repository uses mutex locks for concurrent access
- Demonstrates proper concurrency patterns in Go

### **Repository Pattern**

- Abstract repository interfaces
- Multiple implementations (Memory, MongoDB)
- Easy to swap persistence layers

### **Configuration Pattern**

- Functional options for service configuration
- Flexible dependency injection

### **Domain Modeling**

- Rich domain models with business logic
- Proper aggregate boundaries
- Value objects for data integrity

## 🧪 Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./aggregate
go test ./services
```

## 📚 Learning Resources

This project demonstrates several DDD concepts:

1. **Aggregates**: Customer and Product as aggregate roots
2. **Entities**: Person and Item with identity
3. **Value Objects**: Transaction as immutable data
4. **Repositories**: Data access abstraction
5. **Domain Services**: Business logic orchestration
6. **Application Services**: Use case coordination

## 🔄 Concurrency Demo

The `race_demo.go` file demonstrates:

- Race conditions in concurrent map access
- Proper mutex usage for thread safety
- Go's race detector usage

Run with race detection:

```bash
go run -race race_demo.go
```

## 🛠️ Dependencies

- **UUID Generation**: `github.com/google/uuid`
- **MongoDB Driver**: `go.mongodb.org/mongo-driver/v2`
- **Concurrency**: `golang.org/x/sync`

## 📄 License

This project is for educational purposes, demonstrating DDD patterns in Go.

## 🤝 Contributing

This is a learning project. Feel free to:

- Report issues
- Suggest improvements
- Add new DDD patterns
- Improve documentation

---

**Note**: This is an educational DDD implementation. For production use, consider additional patterns like CQRS, Event Sourcing, and proper error handling strategies.
