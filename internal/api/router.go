package api

import (
	"clinic-system/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// Security fix
	router.SetTrustedProxies(nil)

	// Health route
	router.GET("/health", handlers.HealthCheck)

	return router
}
