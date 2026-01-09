# Quick Start Guide

Get the Go API Boilerplate up and running in minutes!

## Prerequisites

- [Go 1.22+](https://golang.org/doc/install)
- [PostgreSQL 12+](https://www.postgresql.org/download/)
- [Docker](https://docs.docker.com/get-docker/) (optional, for containerized setup)
- [golang-migrate](https://github.com/golang-migrate/migrate) (for migrations)

## Option 1: Quick Start with Docker (Recommended)

The easiest way to get started is using Docker Compose:

### 1. Clone the Repository

```bash
git clone https://github.com/niksbanna/go-api-boilerplate.git
cd go-api-boilerplate
```

### 2. Start Services

```bash
docker-compose up -d
```

This will start:
- PostgreSQL database on port 5432
- API server on port 3000

### 3. Run Migrations

```bash
# Install golang-migrate if not already installed
# On macOS
brew install golang-migrate

# On Linux
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/

# Run migrations
export DB_URL="postgres://postgres:postgres@localhost:5432/api_db?sslmode=disable"
make migrate-up
```

### 4. Test the API

```bash
# Health check
curl http://localhost:3000/health

# Create a user
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# Get all users
curl http://localhost:3000/api/v1/users
```

## Option 2: Local Development Setup

### 1. Clone and Install Dependencies

```bash
git clone https://github.com/niksbanna/go-api-boilerplate.git
cd go-api-boilerplate
go mod download
```

### 2. Setup PostgreSQL

Create a database:

```bash
# Using psql
psql -U postgres
CREATE DATABASE api_db;
\q
```

### 3. Configure Environment

```bash
cp .env.example .env
```

Edit `.env` with your database credentials:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=api_db
```

### 4. Run Migrations

```bash
export DB_URL="postgres://postgres:your_password@localhost:5432/api_db?sslmode=disable"
make migrate-up
```

### 5. Run the Application

```bash
make run
# or
go run cmd/api/main.go
```

The API will be available at http://localhost:3000

## Verify Installation

### Health Check

```bash
curl http://localhost:3000/health
```

Expected response:
```json
{
  "status": "ok",
  "database": "healthy"
}
```

### Create Your First User

```bash
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

Expected response:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-09T12:00:00Z",
  "updated_at": "2024-01-09T12:00:00Z"
}
```

### Retrieve All Users

```bash
curl http://localhost:3000/api/v1/users
```

## Next Steps

1. **Explore the API**: Check out [API.md](API.md) for complete API documentation
2. **Review the Code**: Understand the Clean Architecture structure in the [README](README.md)
3. **Add Features**: Learn how to extend the boilerplate with new entities
4. **Run Tests**: Execute `make test` to run the test suite
5. **Deploy**: Use the Dockerfile for containerized deployments

## Common Commands

```bash
# Build the application
make build

# Run tests
make test

# Run linter
make lint

# Format code
make fmt

# Create a new migration
make migrate-create name=add_products_table

# Stop Docker services
docker-compose down
```

## Troubleshooting

### Database Connection Issues

If you see database connection errors:

1. Verify PostgreSQL is running:
   ```bash
   # Docker
   docker-compose ps
   
   # Local
   pg_isready
   ```

2. Check your `.env` file has correct database credentials

3. Ensure the database exists:
   ```bash
   psql -U postgres -l
   ```

### Port Already in Use

If port 3000 is already in use:

1. Change `SERVER_PORT` in `.env`
2. Update the port in `docker-compose.yml` if using Docker

### Migration Errors

If migrations fail:

1. Check database connection
2. Verify `DB_URL` environment variable is set correctly
3. Run `make migrate-down` to rollback, then `make migrate-up` again

## Getting Help

- Check the [README](README.md) for detailed documentation
- Review [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines
- Open an issue on GitHub for bugs or questions

## What's Next?

Now that you have the boilerplate running, you can:

- Add authentication (JWT, OAuth)
- Implement additional entities (products, orders, etc.)
- Add validation middleware
- Integrate logging and monitoring
- Set up rate limiting
- Add caching with Redis
- Implement pagination
- Add more comprehensive tests

Happy coding! 🚀
