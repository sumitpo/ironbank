package api

import (
	"bank-backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Public routes
	public := r.Group("/api")
	{
		public.GET("/health", HealthCheck) // Add this line
		public.GET("/live", LivenessCheck)
		public.GET("/ready", ReadinessCheck)
		public.POST("/register", Register)
		public.POST("/login", Login)
	}

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		// Account routes
		protected.GET("/accounts", GetAccounts)
		protected.POST("/accounts", CreateAccount)

		// Transaction routes
		protected.POST("/transfer", TransferFunds)
		protected.GET("/accounts/:accountId/transactions", GetTransactions)
	}
}
