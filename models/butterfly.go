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
	ID          uint             `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
	DeletedAt   gorm.DeletedAt   `json:"-" gorm:"index"`
	Scientific  string           `json:"scientific" gorm:"unique,not null"`
	Described   string           `json:"described" gorm:"not null"`
	Family      string           `json:"family" gorm:"not null,default:''"`
	UserId      uint             `json:"user_id" gorm:"not null"`
	Rarity      ButterflyRarity  `json:"rarity" gorm:"type:enum_butterfly_rarity;not null"`
	Daytime     ButterflyDaytime `json:"daytime" gorm:"type:enum_butterfly_daytime;not null"`
	Group       utils.NullString `json:"group"`
	Appearances uint             `json:"appearances" gorm:"default:0;not null"`
	Size        utils.NullString `json:"size"`
	Image       utils.NullString `json:"image"`
	// External models
	User    User                   `json:"user"`
	Details []ButterflyDescription `json:"details" gorm:"ForeignKey:ButterflyId"`
}
