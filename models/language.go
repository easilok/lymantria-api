package models

import (
	"gorm.io/gorm"
)

type Language struct {
	gorm.Model        // include id, created_at, updated_at and deleted_at
	UserId     uint   `json:"user_id" gorm:"not null"`
	User       User   `json:"user"`
	Name       string `json:"name" gorm:"index:idx_name_slug,unique;not null"`
	Slug       string `json:"slug" gorm:"index:idx_name_slug,unique;not null"`
}
