package api

import (
	"clinic-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/health", handlers.HealthCheck)

	return router
}
