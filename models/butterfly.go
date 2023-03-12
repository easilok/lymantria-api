package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type ButterflyRarity string

const (
	COMMON    ButterflyRarity = "common"
	RARE      ButterflyRarity = "rare"
	ULTRARARE ButterflyRarity = "ultrarare"
)

type ButterflyDaytime string

const (
	DAY   ButterflyDaytime = "day"
	NIGHT ButterflyDaytime = "night"
)

type Butterfly struct {
	gorm.Model                   // include id, created_at, updated_at and deleted_at
	UserId      uint             `json:"user_id" gorm:"not null"`
	User        User             `json:"user"`
	Described   sql.NullString   `json:"described" gorm:"not null"`
	Rarity      ButterflyRarity  `json:"rarity" gorm:"type:enum_butterfly_rarity;not null"`
	Daytime     ButterflyDaytime `json:"daytime" gorm:"type:enum_butterfly_daytime;not null"`
	Group       sql.NullString   `json:"group"`
	Appearances uint             `json:"appearances" gorm:"default:0;not null"`
	Size        sql.NullInt32    `json:"size"`
	Image       sql.NullString   `json:"image"`
}
