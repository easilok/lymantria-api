package models

import (
	"time"

	"gorm.io/gorm"
)

type ButterflyAppearances struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
	MonitoringId uint           `json:"monitoring_id" gorm:"not null;index:idx_butterfly_monitoring,unique"`
	ButterflyId  uint           `json:"butterfly_id" gorm:"not null;index:idx_butterfly_monitoring,unique"`
	Quantity     uint           `json:"quantity" gorm:"not null;default:1"`
	Image        string         `json:"image"`
	Observation  string         `j́son:"obs"`
	RegisteredBy uint           `json:"registered_by" gorm:"not null"`
	// Foreign models
	Register  User      `json:"user" gorm:"foreignKey:RegisteredBy"`
	Butterfly Butterfly `json:"butterfly"`
}
