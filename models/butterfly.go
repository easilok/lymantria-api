package models

import (
	"time"

	"github.com/easilok/lymantria-api/utils"
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
	ID          uint                   `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time              `json:"-"`
	UpdatedAt   time.Time              `json:"-"`
	DeletedAt   gorm.DeletedAt         `json:"-" gorm:"index"`
	UserId      uint                   `json:"user_id" gorm:"not null"`
	User        User                   `json:"user"`
	Described   string                 `json:"described" gorm:"not null"`
	Rarity      ButterflyRarity        `json:"rarity" gorm:"type:enum_butterfly_rarity;not null"`
	Daytime     ButterflyDaytime       `json:"daytime" gorm:"type:enum_butterfly_daytime;not null"`
	Group       utils.NullString       `json:"group"`
	Appearances uint                   `json:"appearances" gorm:"default:0;not null"`
	Size        utils.NullInt32        `json:"size"`
	Image       utils.NullString       `json:"image"`
	Details     []ButterflyDescription `json:"details" gorm:"ForeignKey:ButterflyId"`
}
