# Go API Boilerplate

A production-grade boilerplate for building RESTful APIs in Go using the [Fiber](https://gofiber.io/) framework with Clean Architecture principles.

## Features

- ✨ **Clean Architecture**: Separation of concerns with handler, service, and repository layers
- 🚀 **Fiber Framework**: Fast and lightweight web framework
- 🔧 **Configuration Management**: Environment-based configuration with `.env` support
- 🗄️ **Database Integration**: PostgreSQL with connection pooling
- 🔄 **Database Migrations**: Schema versioning with golang-migrate
- 🧪 **Testing Ready**: Structured for unit and integration tests
- 📝 **Linting & Formatting**: Pre-configured golangci-lint
- 🔄 **CI/CD**: GitHub Actions workflows for automated testing and building
- 📦 **Dependency Management**: Go modules
- 🏗️ **Scalable Structure**: Easy to extend and maintain

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── database/
│   │   └── database.go          # Database connection
│   ├── handler/
│   │   ├── health_handler.go   # Health check endpoints
│   │   └── user_handler.go     # User API handlers
│   ├── model/
│   │   └── user.go              # Data models
│   ├── repository/
│   │   └── user_repository.go  # Data access layer
│   └── service/
│       └── user_service.go      # Business logic layer
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   └── 000001_create_users_table.down.sql
├── pkg/
│   └── utils/                   # Shared utilities
├── .github/
│   └── workflows/
│       └── ci.yml               # CI/CD pipeline
├── .env.example                 # Environment variables template
├── .gitignore                   # Git ignore rules
├── .golangci.yml                # Linter configuration
├── go.mod                       # Go module definition
├── Makefile                     # Common tasks automation
└── README.md                    # This file
```

## Prerequisites

- Go 1.22 or higher
- PostgreSQL 12 or higher
- golang-migrate (for database migrations)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/niksbanna/go-api-boilerplate.git
cd go-api-boilerplate
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Configure Environment

Copy the example environment file and update with your values:

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```env
SERVER_PORT=3000
SERVER_HOST=0.0.0.0
APP_ENV=development

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=api_db
DB_SSLMODE=disable
```

### 4. Setup Database

Create a PostgreSQL database:

```bash
createdb api_db
```

Run migrations:

```bash
# Set database URL
export DB_URL="postgres://postgres:your_password@localhost:5432/api_db?sslmode=disable"

# Run migrations
make migrate-up
```

### 5. Run the Application

```bash
# Using Make
make run

# Or directly with Go
go run cmd/api/main.go
```

The API will be available at `http://localhost:3000`

## API Endpoints

### Health Check

```bash
GET /health
```

### Users API

```bash
# Get all users
GET /api/v1/users

# Get user by ID
GET /api/v1/users/:id

# Create user
POST /api/v1/users
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com"
}

# Update user
PUT /api/v1/users/:id
Content-Type: application/json

{
  "name": "Jane Doe",
  "email": "jane@example.com"
}

# Delete user
DELETE /api/v1/users/:id
```

## Development

### Building

```bash
# Build the application
make build

# The binary will be in bin/api
./bin/api
```

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Linting

```bash
# Run linter
make lint
```

### Formatting

```bash
# Format code
make fmt
```

### Database Migrations

```bash
# Create a new migration
make migrate-create name=add_new_table

# Run migrations up
make migrate-up

# Rollback migrations
make migrate-down
```

## Clean Architecture

This boilerplate follows Clean Architecture principles with clear separation of concerns:

### Layers

1. **Handler Layer** (`internal/handler/`)
   - Handles HTTP requests and responses
   - Validates input
   - Calls service layer
   - Returns appropriate HTTP status codes

2. **Service Layer** (`internal/service/`)
   - Contains business logic
   - Validates business rules
   - Orchestrates data flow
   - Independent of HTTP and database details

3. **Repository Layer** (`internal/repository/`)
   - Handles database operations
   - Implements data access interfaces
   - Isolates database logic

### Benefits

- **Testability**: Each layer can be tested independently
- **Maintainability**: Changes in one layer don't affect others
- **Scalability**: Easy to add new features
- **Flexibility**: Easy to swap implementations (e.g., change database)

## CI/CD

The project includes GitHub Actions workflows for:

- **Linting**: Ensures code quality with golangci-lint
- **Formatting**: Checks code formatting with gofmt
- **Building**: Compiles the application
- **Testing**: Runs tests with PostgreSQL service

Workflows run automatically on push and pull requests to `main` and `develop` branches.

## Adding New Features

### Adding a New Entity

1. Create the model in `internal/model/`
2. Create the repository interface and implementation in `internal/repository/`
3. Create the service interface and implementation in `internal/service/`
4. Create the handler in `internal/handler/`
5. Register routes in `cmd/api/main.go`
6. Create migrations for database schema

### Example: Adding a Product Entity

```go
// 1. Model (internal/model/product.go)
type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

// 2. Repository interface (internal/repository/product_repository.go)
type ProductRepository interface {
    Create(ctx context.Context, product *model.Product) error
    GetByID(ctx context.Context, id int) (*model.Product, error)
}

// 3. Service (internal/service/product_service.go)
type ProductService interface {
    CreateProduct(ctx context.Context, req *model.CreateProductRequest) (*model.Product, error)
}

// 4. Handler (internal/handler/product_handler.go)
type ProductHandler struct {
    productService service.ProductService
}

// 5. Register routes (cmd/api/main.go)
products := api.Group("/products")
products.Post("/", productHandler.CreateProduct)
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_PORT` | Server port | `3000` |
| `SERVER_HOST` | Server host | `0.0.0.0` |
| `APP_ENV` | Application environment | `development` |
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | `api_db` |
| `DB_SSLMODE` | Database SSL mode | `disable` |

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License.

## Support

For questions and support, please open an issue in the GitHub repository.