package main

import (
	"github.com/easilok/lymantria-api/config"
	c "github.com/easilok/lymantria-api/controllers"
	"github.com/easilok/lymantria-api/database"
	"github.com/easilok/lymantria-api/helpers"
	"github.com/easilok/lymantria-api/middlewares"
	"github.com/easilok/lymantria-api/routes"
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

var db *gorm.DB

func main() {
	// loadConfig()
	config.SetupLogger()

	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	db = database.ConnectDatabase(&databaseConfig, true)
	controllers := c.NewBaseHandler(db)

	helpers.CheckTokenSecrets()

	routes.ApiRoutes(r, controllers)
	routes.AuthRoutes(r, controllers)

	r.Run("0.0.0.0:8080")

}
