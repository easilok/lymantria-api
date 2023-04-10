package models

import (
	"time"

	"github.com/easilok/lymantria-api/utils"
	"gorm.io/gorm"
)

type Monitoring struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	RegisteredAt  time.Time      `json:"registered_at" gorm:"not null"`
	Local         string         `json:"local" gorm:"not null"`
	Name          string         `json:"name" gorm:"not null"`
	Longitude     float32        `json:"longitude"`
	Latitude      float32        `json:"latitude"`
	Observation   string         `j́son:"obs"`
	HostedBy      uint           `json:"hosted_by" gorm:"not null"`
	TimestampEnd  utils.NullTime `json:"timestamp_end"`
	Temperature   float32        `json:"temperature" gorm:"not null,default:0"`
	Humidity      float32        `json:"humidity" gorm:"not null,default:0"`
	Wind          string         `json:"wind" gorm:"not null,default:''"`
	Precipitation float32        `json:"precipitation" gorm:"not null,default:0"`
	Sky           string         `json:"sky" gorm:"not null,default:''"`
	// External models
	Host        User                   `json:"user" gorm:"foreignKey:HostedBy"`
	Appearances []ButterflyAppearances `json:"appearances" gorm:"ForeignKey:MonitoringId"`
}
