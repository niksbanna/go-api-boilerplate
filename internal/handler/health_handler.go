package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/niksbanna/go-api-boilerplate/internal/database"
)

// HealthHandler handles health check requests
type HealthHandler struct {
	db *database.Database
}

// NewHealthHandler creates a new health handler
func NewHealthHandler(db *database.Database) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

// Health handles GET /health
func (h *HealthHandler) Health(c *fiber.Ctx) error {
	// Check database connection
	dbStatus := "healthy"
	if err := h.db.Health(); err != nil {
		dbStatus = "unhealthy"
	}

	return c.JSON(fiber.Map{
		"status":   "ok",
		"database": dbStatus,
	})
}
