package entity

import (
	"time"

	"gorm.io/gorm"
)

type Place struct {
	gorm.Model
	Locate       string
	BookInfolist []BookInfolist `gorm:"foreignKey:PlaceID"`
}

type TimePeriod struct {
	gorm.Model
	Period       string
	BookInfolist []BookInfolist `gorm:"foreignKey:TimePeriodID"`
}

type BookInfolist struct {
	gorm.Model
	ServiceID  *uint
	Service    Service `gorm:"references:ID"`
	PlaceID    *uint
	Place      Place `gorm:"references:ID"`
	TimePeriodID     *uint
	TimePeriod TimePeriod `gorm:"references:ID"`
	MemberID   *uint
	Member     Member `gorm:"references:ID"`
	BDate      time.Time
}