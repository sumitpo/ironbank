package api

import (
	"bank-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthStatus struct {
	Status  string `json:"status"`
	DB      string `json:"db,omitempty"`
	Version string `json:"version"`
}

// HealthCheck responds with API health status
// @Summary Show API health status
// @Description get API health status
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthStatus
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	status := HealthStatus{
		Status:  "healthy",
		Version: "1.0.0",
	}

	// Check database connection
	if err := models.GetDB().Exec("SELECT 1").Error; err != nil {
		status.Status = "unhealthy"
		status.DB = err.Error()
		c.JSON(http.StatusServiceUnavailable, status)
		return
	}

	status.DB = "connected"
	c.JSON(http.StatusOK, status)
}

// In api/health.go
func LivenessCheck(c *gin.Context) {

	status := HealthStatus{
		Status:  "live",
		Version: "1.0.0",
		DB:      "not known",
	}
	c.JSON(http.StatusOK, status)
}

func ReadinessCheck(c *gin.Context) {
	if err := models.CheckDBHealth(); err != nil {
		c.Status(http.StatusServiceUnavailable)
		return
	}
	c.Status(http.StatusOK)
}
