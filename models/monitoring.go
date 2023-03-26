package models

import (
	"time"

	"gorm.io/gorm"
)

type Monitoring struct {
	ID           uint                   `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time              `json:"-"`
	UpdatedAt    time.Time              `json:"-"`
	DeletedAt    gorm.DeletedAt         `json:"-" gorm:"index"`
	RegisteredAt time.Time              `json:"registered_at" gorm:"not null"`
	Local        string                 `json:"local" gorm:"not null"`
	Name         string                 `json:"name" gorm:"not null"`
	Longitude    uint                   `json:"longitude"`
	Latitude     uint                   `json:"latitude"`
	Observation  string                 `j́son:"obs"`
	HostedBy     uint                   `json:"hosted_by" gorm:"not null"`
	Host         User                   `json:"user" gorm:"foreignKey:HostedBy"`
	Appearances  []ButterflyAppearances `json:"appearances" gorm:"ForeignKey:MonitoringId"`
}
