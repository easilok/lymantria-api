package routes

import (
	"github.com/easilok/lymantria-api/controllers"
	"github.com/easilok/lymantria-api/middlewares"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine, controllers *controllers.BaseHandler) {
	apiGroup := r.Group("/api")
	// Butterfly routes
	apiGroup.GET("/butterfly", controllers.GetAllButterflies)
	// apiGroup.POST("/butterfly", TokenAuthMiddleware(), controllers.CreateButterfly)
	apiGroup.PUT("/butterfly/:butterflyId", middlewares.TokenAuthMiddleware(), controllers.UpdateButterfly)
	apiGroup.DELETE("/butterfly/:butterflyId", middlewares.TokenAuthMiddleware(), controllers.DeleteButterfly)

	// Monitoring routes
	apiGroup.GET("/monitoring", controllers.GetAllMonitorings)
	apiGroup.GET("/monitoring/latest", controllers.GetMonitoringLatest)
	apiGroup.GET("/monitoring/:monitoringId", controllers.GetMonitoring)
	// apiGroup.POST("/monitoring", TokenAuthMiddleware(), controllers.CreateMonitoring)
	// apiGroup.PUT("/monitoring/:monitoringId", TokenAuthMiddleware(), controllers.UpdateMonitoring)
	// apiGroup.DELETE("/monitoring/:monitoringId", TokenAuthMiddleware(), controllers.DeleteMonitoring)
}
