package models

import (
	"time"

	"github.com/easilok/lymantria-api/utils"
	"gorm.io/gorm"
)

type Monitoring struct {
	ID                 uint              `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time         `json:"-"`
	UpdatedAt          time.Time         `json:"-"`
	DeletedAt          gorm.DeletedAt    `json:"-" gorm:"index"`
	RegisteredAt       time.Time         `json:"registered_at" gorm:"not null"`
	Local              string            `json:"local" gorm:"not null"`
	Name               string            `json:"name" gorm:"not null"`
	Longitude          float32           `json:"longitude"`
	Latitude           float32           `json:"latitude"`
	Observation        string            `j́son:"obs"`
	HostedBy           uint              `json:"hosted_by" gorm:"not null"`
	TimestampEnd       utils.NullTime    `json:"timestamp_end"`
	TemperatureStart   float64           `json:"temperature_start" gorm:"not null,default:0"`
	HumidityStart      float64           `json:"humidity_start" gorm:"not null,default:0"`
	WindStart          string            `json:"wind_start" gorm:"not null,default:''"`
	PrecipitationStart float64           `json:"precipitation_start" gorm:"not null,default:0"`
	SkyStart           string            `json:"sky_start" gorm:"not null,default:''"`
	TemperatureEnd     utils.NullFloat64 `json:"temperature_end" `
	HumidityEnd        utils.NullFloat64 `json:"humidity_end" `
	WindEnd            utils.NullString  `json:"wind_end" "`
	PrecipitationEnd   utils.NullFloat64 `json:"precipitation_end" `
	SkyEnd             utils.NullString  `json:"sky_end" "`
	// External models
	Host        User                   `json:"user" gorm:"foreignKey:HostedBy"`
	Appearances []ButterflyAppearances `json:"appearances" gorm:"ForeignKey:MonitoringId"`
}
