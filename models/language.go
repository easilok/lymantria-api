package models

import (
	"time"

	"gorm.io/gorm"
)

type Language struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	UserId    uint           `json:"user_id" gorm:"not null"`
	User      User           `json:"user"`
	Name      string         `json:"name" gorm:"index:idx_name_slug,unique;not null"`
	Slug      string         `json:"slug" gorm:"index:idx_name_slug,unique;not null"`
}
