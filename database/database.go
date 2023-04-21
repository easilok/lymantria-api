package database

import (
	"fmt"

	"github.com/easilok/lymantria-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Hostname string
	Username string
	Password string
	Database string
	Port     uint16
}

func FirstSetup(db *gorm.DB) {
	var user models.User
	hashedPassword, err := models.HashPassword("admin")
	if err != nil {
		return
	}
	if err := db.First(&user).Error; err != nil {
		user.Email = "admin@lymantria.com"
		user.Name = "Lymantria"
		user.Password = hashedPassword
		user.Permissions = []string{"admin"}
		fmt.Println(user)
		db.Save(&user)
	}
}

func ConnectDatabase(config *DatabaseConfig, firstSetup bool) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Lisbon",
		config.Hostname,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Language{})
	database.AutoMigrate(&models.Butterfly{})
	database.AutoMigrate(&models.ButterflyDescription{})
	database.AutoMigrate(&models.Monitoring{})
	database.AutoMigrate(&models.ButterflyAppearances{})

	if firstSetup {
		FirstSetup(database)
	}

	return database
}
