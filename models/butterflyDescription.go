package models

import (
	"time"

	"gorm.io/gorm"
)

type ButterflyDescription struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	LanguageId  uint           `json:"language_id" gorm:"index:idx_butterfly_description_language,unique;not null"`
	ButterflyId uint           `json:"butterfly_id" gorm:"index:idx_butterfly_description_language,unique;not null"`
	UserId      uint           `json:"user_id" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	CommonName  string         `json:"common_name"`
	Observation string         `json:"observation"`
	// Foreign models
	User     User `json:"user"`
	Language Language
	// Butterfly Butterfly
}
