package main

import (
	"net/http"

	config "github.com/easilok/lymantria-api/config"
	c "github.com/easilok/lymantria-api/controllers"
	"github.com/easilok/lymantria-api/database"
	"github.com/easilok/lymantria-api/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var databaseConfig database.DatabaseConfig = database.DatabaseConfig{
	Hostname: "localhost",
	Username: "lymantria",
	Password: "lymantria-pass",
	Database: "lymantria",
	Port:     5432,
}

// TODO - This should be removed after token middlware implementation
var db *gorm.DB

func main() {
	// loadConfig()
	config.SetupLogger()

	r := gin.Default()
	r.Use(CORSMiddleware())

	db = database.ConnectDatabase(&databaseConfig, true)
	controllers := c.NewBaseHandler(db)

	// helpers.CheckTokenSecrets()

	apiGroup := r.Group("/api")
	{
		// Butterfly routes
		apiGroup.GET("/butterfly", controllers.GetAllButterflies)
		// apiGroup.POST("/butterfly", TokenAuthMiddleware(), controllers.CreateButterfly)
		apiGroup.PUT("/butterfly/:butterflyId", TokenAuthMiddleware(), controllers.UpdateButterfly)
		apiGroup.DELETE("/butterfly/:butterflyId", TokenAuthMiddleware(), controllers.DeleteButterfly)

		// Monitoring routes
		apiGroup.GET("/monitoring", controllers.GetAllMonitorings)
		apiGroup.GET("/monitoring/latest", controllers.GetMonitoringLatest)
		apiGroup.GET("/monitoring/:monitoringId", controllers.GetMonitoring)
		// apiGroup.POST("/monitoring", TokenAuthMiddleware(), controllers.CreateMonitoring)
		// apiGroup.PUT("/monitoring/:monitoringId", TokenAuthMiddleware(), controllers.UpdateMonitoring)
		// apiGroup.DELETE("/monitoring/:monitoringId", TokenAuthMiddleware(), controllers.DeleteMonitoring)
	}

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
		authGroup.GET("/logout", TokenAuthMiddleware(), controllers.Logout)
		authGroup.POST("/refresh", controllers.Refresh)
		// authGroup.POST("/register", TokenAuthMiddleware(), controllers.Register)
		authGroup.PATCH("/password", TokenAuthMiddleware(), controllers.Password)
	}

	r.Run("0.0.0.0:8080")

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		au, err := helpers.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Set("userId", au.UserId)
		c.Next()
	}
}
