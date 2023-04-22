package routes

import (
	"github.com/easilok/lymantria-api/controllers"
	"github.com/easilok/lymantria-api/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, controllers *controllers.BaseHandler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.GET("/logout", middlewares.TokenAuthMiddleware(), controllers.Logout)
		authGroup.POST("/refresh", controllers.Refresh)
		// authGroup.POST("/register", TokenAuthMiddleware(), controllers.Register)
		authGroup.PATCH("/password", middlewares.TokenAuthMiddleware(), controllers.Password)
	}
}
