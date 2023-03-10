package main

import (
	c "github.com/easilok/lymantria-api/controller"
	"github.com/easilok/lymantria-api/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

var databaseConfig database.DatabaseConfig = database.DatabaseConfig{
	Hostname: "localhost",
	Username: "lymantria",
	Password: "lymantria-pass",
	Database: "lymantria",
	Port:     5432,
}

func main() {
	// loadConfig()
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})
	// If you wish to add the calling method as a field, instruct the logger via:
	// log.SetReportCaller(true)

	r := gin.Default()
	r.Use(CORSMiddleware())

	db := database.ConnectDatabase(&databaseConfig)
	controllers := c.NewBaseHandler(db)

	// helpers.CheckTokenSecrets()

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/butterfly", TokenAuthMiddleware(), controllers.GetButterfly)
		// apiGroup.POST("/butterfly", TokenAuthMiddleware(), controllers.FavoriteNote)
		// apiGroup.PUT("/butterfly/:butterfly", TokenAuthMiddleware(), controllers.UpdateNote)
		// apiGroup.DELETE("/butterfly/:butterfly", TokenAuthMiddleware(), controllers.DeleteNote)
	}

	// authGroup := r.Group("/auth")
	// {
	// 	authGroup.POST("/login", controllers.Login)
	// 	authGroup.GET("/logout", TokenAuthMiddleware(), controllers.Logout)
	// 	authGroup.POST("/refresh", controllers.Refresh)
	// 	authGroup.POST("/register", TokenAuthMiddleware(), controllers.Register)
	// 	authGroup.PATCH("/password", TokenAuthMiddleware(), controllers.Password)
	// }

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
		// au, err := helpers.ExtractTokenMetadata(c.Request)
		// // err := helpers.TokenValid(c.Request)
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, err.Error())
		// 	c.Abort()
		// 	return
		// }
		// c.Set("userId", au.UserId)
		c.Next()
	}
}
