package models

import (
	"gorm.io/gorm"
)

type ButterflyAppearances struct {
	gorm.Model          // include id, created_at, updated_at and deleted_at
	MonitoringId uint   `json:"monitoring_id" gorm:"not null;index:idx_butterfly_monitoring,unique"`
	ButterflyId  uint   `json:"butterfly_id" gorm:"not null;index:idx_butterfly_monitoring,unique"`
	Quantity     uint   `json:"quantity" gorm:"not null;default:1"`
	Image        string `json:"image"`
	Observation  string `j́son:"obs"`
	RegisteredBy uint   `json:"registered_by" gorm:"not null"`
	// Foreign models
	User       User `json:"user" gorm:"foreignKey:RegisteredBy"`
	Monitoring Monitoring
	Butterfly  Butterfly
}
