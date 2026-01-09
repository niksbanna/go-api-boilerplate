# API Documentation

## Base URL

```
http://localhost:3000
```

## Endpoints

### Health Check

Check the health status of the API and database connection.

**Endpoint:** `GET /health`

**Response:**

```json
{
  "status": "ok",
  "database": "healthy"
}
```

**Status Codes:**
- `200 OK`: Service is healthy

---

### Users API

#### Get All Users

Retrieve a list of all users.

**Endpoint:** `GET /api/v1/users`

**Response:**

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2024-01-09T12:00:00Z",
    "updated_at": "2024-01-09T12:00:00Z"
  }
]
```

**Status Codes:**
- `200 OK`: Successfully retrieved users
- `500 Internal Server Error`: Server error

---

#### Get User by ID

Retrieve a specific user by their ID.

**Endpoint:** `GET /api/v1/users/:id`

**Path Parameters:**
- `id` (integer): User ID

**Response:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-09T12:00:00Z",
  "updated_at": "2024-01-09T12:00:00Z"
}
```

**Status Codes:**
- `200 OK`: Successfully retrieved user
- `400 Bad Request`: Invalid user ID
- `404 Not Found`: User not found
- `500 Internal Server Error`: Server error

---

#### Create User

Create a new user.

**Endpoint:** `POST /api/v1/users`

**Request Headers:**
- `Content-Type: application/json`

**Request Body:**

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "created_at": "2024-01-09T12:00:00Z",
  "updated_at": "2024-01-09T12:00:00Z"
}
```

**Status Codes:**
- `201 Created`: User successfully created
- `400 Bad Request`: Invalid request body
- `500 Internal Server Error`: Server error

**Validation:**
- `name`: Required, non-empty string
- `email`: Required, valid email format

---

#### Update User

Update an existing user.

**Endpoint:** `PUT /api/v1/users/:id`

**Path Parameters:**
- `id` (integer): User ID

**Request Headers:**
- `Content-Type: application/json`

**Request Body:**

```json
{
  "name": "Jane Doe",
  "email": "jane@example.com"
}
```

**Note:** All fields are optional. Only provided fields will be updated.

**Response:**

```json
{
  "id": 1,
  "name": "Jane Doe",
  "email": "jane@example.com",
  "created_at": "2024-01-09T12:00:00Z",
  "updated_at": "2024-01-09T12:30:00Z"
}
```

**Status Codes:**
- `200 OK`: User successfully updated
- `400 Bad Request`: Invalid user ID or request body
- `404 Not Found`: User not found
- `500 Internal Server Error`: Server error

---

#### Delete User

Delete a user by ID.

**Endpoint:** `DELETE /api/v1/users/:id`

**Path Parameters:**
- `id` (integer): User ID

**Response:**

No content (empty response body)

**Status Codes:**
- `204 No Content`: User successfully deleted
- `400 Bad Request`: Invalid user ID
- `404 Not Found`: User not found
- `500 Internal Server Error`: Server error

---

## Error Response Format

All error responses follow this format:

```json
{
  "error": "Error message description"
}
```

## Testing the API

### Using cURL

```bash
# Health check
curl http://localhost:3000/health

# Get all users
curl http://localhost:3000/api/v1/users

# Get user by ID
curl http://localhost:3000/api/v1/users/1

# Create user
curl -X POST http://localhost:3000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# Update user
curl -X PUT http://localhost:3000/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","email":"jane@example.com"}'

# Delete user
curl -X DELETE http://localhost:3000/api/v1/users/1
```

### Using Postman

Import the `postman_collection.json` file into Postman to test all endpoints.

## Rate Limiting

Currently, there is no rate limiting implemented. Consider adding rate limiting middleware for production use.

## Authentication

Currently, there is no authentication implemented. This is a boilerplate project. For production use, consider adding:

- JWT authentication
- API key authentication
- OAuth 2.0

## CORS

CORS is enabled for all origins in development mode. For production, configure appropriate CORS settings in `cmd/api/main.go`.
