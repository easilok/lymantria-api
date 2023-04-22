package test

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/easilok/lymantria-api/database"
	"github.com/easilok/lymantria-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var databaseConfig database.DatabaseConfig = database.DatabaseConfig{
	Hostname: "localhost",
	Username: "test",
	Password: "test-pass",
	Database: "lymantria_test",
	Port:     5434,
}

func CreateExampleUser(db *gorm.DB) models.User {
	var user models.User
	if err := db.First(&user).Error; err != nil {
		user.Email = "test@test.com"
		user.Name = "test"
		hashedPassword, _ := models.HashPassword("123456")
		user.Password = hashedPassword
		db.Save(&user)
	}
	return user
}

func ConnectTestDatabase() *gorm.DB {
	db := database.ConnectDatabase(&databaseConfig, false)

	return db
}

func MockJsonPost(c *gin.Context /* the test context */, content interface{}, method string) error {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		return err
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))

	return nil
}
