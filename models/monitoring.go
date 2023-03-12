package models

import (
	"time"

	"gorm.io/gorm"
)

type Monitoring struct {
	gorm.Model             // include id, created_at, updated_at and deleted_at
	RegisteredAt time.Time `json:"registered_at" gorm:"not null"`
	Local        string    `json:"local" gorm:"not null"`
	Name         string    `json:"name" gorm:"not null"`
	Longitude    uint      `json:"longitude"`
	Latitude     uint      `json:"latitude"`
	Observation  string    `j́son:"obs"`
	HostedBy     uint      `json:"hosted_by" gorm:"not null"`
	User         User      `json:"user" gorm:"foreignKey:HostedBy"`
}
