package entity

import (
	"time"

	"gorm.io/gorm"
)

type EquipmentName struct {
	gorm.Model

	Name          string
	EquipmentList []EquipmentList `gorm:"foreignKey:EquipmentNameID"`
	Notify        []Notify        `gorm:"foreignKey:EquipmentNameID"`
}

type RunNumber struct {
	gorm.Model

	Number        string
	EquipmentList []EquipmentList `gorm:"foreignKey:RunNumberID"`
	Notify        []Notify        `gorm:"foreignKey:RunNumberID"`
}

type EquipmentList struct {
	gorm.Model

	EmployeeID *uint
	Employee   Employee 

	EquipmentNameID *uint
	EquipmentName   EquipmentName 

	RunNumberID *uint
	RunNumber   RunNumber 

	EquipmentBookingList []EquipmentBookingList `gorm:"foreignKey:EquipmentListID"`
	dateTime             time.Time
}
