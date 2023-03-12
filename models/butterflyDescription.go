package models

import (
	"gorm.io/gorm"
)

type ButterflyDescription struct {
	gorm.Model         // include id, created_at, updated_at and deleted_at
	LanguageId  uint   `json:"language_id" gorm:"index:idx_butterfly_description_language,unique;not null"`
	ButterflyId uint   `json:"butterfly_id" gorm:"index:idx_butterfly_description_language,unique;not null"`
	UserId      uint   `json:"user_id" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	CommonName  string `json:"common_name"`
	Observation string `json:"observation"`
	// Foreign models
	User      User `json:"user"`
	Language  Language
	Butterfly Butterfly
}
